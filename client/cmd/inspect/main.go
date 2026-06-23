package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/go-routeros/routeros/v3"
)

func main() {
	if err := run(); err != nil {
		log.Println(fmt.Errorf("application failed: %w", err))
		os.Exit(1)
	}
}

func run() error {
	mikrotik := client.NewClient(client.GetConfigFromEnv())

	mc, err := mikrotik.GetMikrotikClient()
	if err != nil {
		return err
	}

	item, err := inspectConsoleCommand(mc, &InspectConfig{
		Root:  "/ip/dns/static",
		Depth: 3,
	})
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, string(b))

	return nil
}

/*

- run request=child to get list of available commands/arguments/subcommands
- run request=completion on each `arg` item to populate valid values

*/

func inspectConsoleCommand(mc *routeros.Client, config *InspectConfig) (*Item, error) {

	root := strings.TrimRight(config.Root, "/")
	items, err := inspectPath(mc, root, "child")
	if err != nil {
		return nil, err
	}
	var rootNode *Item
	for _, v := range items {
		if v.Type == TypeSelf {
			rootNode = &Item{
				Self:     root,
				Name:     v.Name,
				NodeType: v.NodeType,
			}
			break
		}
	}

	if rootNode == nil {
		return nil, fmt.Errorf("could not get root node definition")
	}

	itemsQueue := []*Item{rootNode}
	var nextItemsBatch []*Item

	for i := 1; len(itemsQueue) > 0 && i < 10; i++ {
		if i > config.Depth {
			return rootNode, nil
		}

		// node := itemsQueue[0]
		// items, err = inspectPath(mc, node.Self, "child")
		// if err != nil {
		// 	return nil, err
		// }

		// RANGE_LOOP:
		for _, node := range itemsQueue {
			// if v.Type == TypeSelf {
			// 	switch v.NodeType {
			// 	case NodeTypeCmd:
			// 		args, err := getCommandArguments(mc, node.Self)
			// 		if err != nil {
			// 			return nil, err
			// 		}
			// 		node.Arguments = args

			// 		// short path - `cmd` should not have children
			// 		break RANGE_LOOP
			// 	}
			// 	continue
			// }

			// var discoveredItem *Item
			switch node.NodeType {
			case NodeTypeCmd:
				args, err := getCommandArguments(mc, node.Self)
				if err != nil {
					return nil, err
				}
				node.Arguments = args
				continue
			case NodeTypeDir:
				children, err := getNodeChildren(mc, node.Self)
				if err != nil {
					return nil, err
				}
				node.Children = append(node.Children, children...)
				nextItemsBatch = append(nextItemsBatch, children...)
			default:
				return nil, fmt.Errorf("unsupported node type: %s", node.NodeType)
			}
		}

		itemsQueue = nextItemsBatch
	}

	return rootNode, nil
}

func inspectPath(c *routeros.Client, commandPath string, request string) ([]ConsoleItem, error) {
	normalizedCommand := commandPath
	pathParam := "input"
	switch request {
	case "child", "completion":
		normalizedCommand = strings.ReplaceAll(commandPath[1:], "/", ",")
		pathParam = "path"
	}
	cmd := []string{"/console/inspect", "as-value", "=" + pathParam + "=" + normalizedCommand, "=request=" + request}
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
	items, err := inspectPath(c, command, "child")
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

		completions, err := inspectPath(c, command+"/"+v.Name, "completion")
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

func getNodeChildren(c *routeros.Client, command string) ([]*Item, error) {
	items, err := inspectPath(c, command, "child")
	if err != nil {
		return nil, err
	}

	var children []*Item
	for _, v := range items {
		if v.Type != TypeChild {
			continue
		}
		children = append(children, &Item{
			Self:     command + "/" + v.Name,
			Name:     v.Name,
			NodeType: v.NodeType,
		})
	}

	return children, nil
}
