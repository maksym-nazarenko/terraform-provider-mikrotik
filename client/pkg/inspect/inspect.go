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
			continue
		case NodeTypeDir, NodeTypePath:
			children, err := getNodeChildren(c, node.Self)
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, children...)
			nextItems = append(nextItems, children...)
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
