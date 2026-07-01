package inspect

import (
	"fmt"
	"strings"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/go-routeros/routeros/v3"
)

const (
	requestChild      inspectRequest = "child"
	requestCompletion inspectRequest = "completion"
)

type (
	inspectRequest string
)

// Do runs inspection against remote RouterOS instance.
func Do(c *routeros.Client, config *Config) (*Node, error) {
	root := strings.TrimRight(config.Root, "/")
	items, err := inspectPath(c, root, requestChild)
	if err != nil {
		return nil, err
	}

	var rootNode *Node
	for _, v := range items {
		if v.Type == TypeSelf {
			rootNode = &Node{
				Self: root,
				Name: v.Name,
				Type: v.NodeType,
			}
			break
		}
	}

	if rootNode == nil {
		return nil, fmt.Errorf("could not get root node definition")
	}

	itemsQueue := []*Node{rootNode}
	var nextItemsBatch []*Node

	for i := 1; len(itemsQueue) > 0 && i < 10; i++ {
		if config.Depth >= 0 && i > config.Depth {
			return rootNode, nil
		}

		nextItemsBatch, err = processNodes(c, itemsQueue)
		if err != nil {
			return nil, err
		}
		itemsQueue = nextItemsBatch
	}

	return rootNode, nil
}

func processNodes(c *routeros.Client, nodes []*Node) ([]*Node, error) {
	var nextItems []*Node

	for _, node := range nodes {
		switch node.Type {
		case NodeTypeCmd:
			args, err := getCommandArguments(c, node.Self)
			if err != nil {
				return nil, err
			}
			node.Arguments = args
			if node.ArgumentsMap == nil {
				node.ArgumentsMap = make(map[string]*Argument)
			}
			for i := range args {
				node.ArgumentsMap[args[i].Name] = args[i]
			}
			continue
		case NodeTypeDir, NodeTypePath:
			children, err := getNodeChildren(c, node.Self)
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, children...)
			nextItems = append(nextItems, children...)
			if node.ChildrenMap == nil {
				node.ChildrenMap = make(map[string]*Node)
			}
			for i := range children {
				node.ChildrenMap[children[i].Name] = children[i]
			}

			if addCmd, ok := node.ChildrenMap["add"]; ok && addCmd.Type == NodeTypeCmd {
				args, err := getCommandArguments(c, addCmd.Self)
				if err != nil {
					return nil, err
				}
				addCmd.Arguments = args
				if addCmd.ArgumentsMap == nil {
					addCmd.ArgumentsMap = make(map[string]*Argument)
				}
				for i := range args {
					addCmd.ArgumentsMap[args[i].Name] = args[i]
				}
				readonlyProperties, err := getNodeReadonlyProperties(c, node)
				if err != nil {
					return nil, err
				}
				node.ReadonlyPropertiesMap = readonlyProperties
			}
		default:
			return nil, fmt.Errorf("unsupported node type: %s", node.Type)
		}
	}

	return nextItems, nil
}

func inspectPath(c *routeros.Client, commandPath string, request inspectRequest) ([]ConsoleItem, error) {
	normalizedCommand := commandPath
	pathParam := "input"
	switch request {
	case requestChild, requestCompletion:
		normalizedCommand = strings.ReplaceAll(commandPath[1:], "/", ",")
		pathParam = "path"
	}
	cmd := []string{"/console/inspect", "as-value", "=" + pathParam + "=" + normalizedCommand, "=request=" + string(request)}
	reply, err := c.RunArgs(cmd)
	if err != nil {
		return nil, err
	}
	var items []ConsoleItem
	if err := client.Unmarshal(*reply, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func getCommandArguments(c *routeros.Client, command string) ([]*Argument, error) {
	items, err := inspectPath(c, command, requestChild)
	if err != nil {
		return nil, err
	}

	var args []*Argument
	for _, v := range items {
		if v.Type != TypeChild || v.NodeType != NodeTypeArg {
			continue
		}
		arg := Argument{
			Name: v.Name,
		}

		completions, err := inspectPath(c, command+"/"+v.Name, requestCompletion)
		if err != nil {
			return nil, err
		}

		for _, compl := range completions {
			if compl.Type != TypeCompletion || !compl.Show {
				continue
			}
			arg.Options = append(arg.Options, compl.Completion)
		}
		args = append(args, &arg)
	}

	return args, nil
}

func getNodeReadonlyProperties(c *routeros.Client, node *Node) (map[string]*Property, error) {
	if node.Type != NodeTypeDir && node.Type != NodeTypePath {
		return nil, nil
	}
	if _, ok := node.ChildrenMap["add"]; !ok {
		return nil, nil
	}

	resourceBasePath := node.Self
	argumentsMap := node.ChildrenMap["add"].ArgumentsMap
	if len(argumentsMap) == 0 {
		return nil, nil
	}

	allProperties, err := inspectPath(c, resourceBasePath+"/print,proplist", requestCompletion)
	if err != nil {
		return nil, err
	}

	readonlyProperties := make(map[string]*Property)
	for _, prop := range allProperties {
		if !prop.Show {
			continue
		}
		if _, ok := argumentsMap[prop.Completion]; ok {
			// property is listed as command argument, hence writable
			continue
		}
		p := &Property{Name: prop.Completion}
		readonlyProperties[p.Name] = p
	}

	return readonlyProperties, nil
}

func getNodeChildren(c *routeros.Client, command string) ([]*Node, error) {
	items, err := inspectPath(c, command, requestChild)
	if err != nil {
		return nil, err
	}

	var children []*Node
	for _, v := range items {
		if v.Type != TypeChild {
			continue
		}
		children = append(children, &Node{
			Self: command + "/" + v.Name,
			Name: v.Name,
			Type: v.NodeType,
		})
	}

	return children, nil
}
