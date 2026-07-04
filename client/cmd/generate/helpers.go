package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/inspect"
)

func readDefinitionFile(path string) (*inspect.Node, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open definition file %q: %w", path, err)
	}
	defer f.Close()

	var rootNode inspect.Node
	if err := json.NewDecoder(f).Decode(&rootNode); err != nil {
		return nil, fmt.Errorf("failed to decode definition file %q: %w", path, err)
	}

	return &rootNode, nil
}

func findSubNode(root *inspect.Node, basePath string) (*inspect.Node, error) {
	if root == nil {
		return nil, fmt.Errorf("root node cannot be nil")
	}

	if len(basePath) < 1 {
		return nil, fmt.Errorf("basePath cannot be empty")
	}

	if root.Self == basePath {
		return root, nil
	}

	subPath := strings.Split(strings.TrimLeft(strings.TrimPrefix(basePath, root.Self), "/"), "/")
	node := root
	for i := range subPath {
		if nextNode, ok := node.ChildrenMap[subPath[i]]; ok {
			node = nextNode
			continue
		}

		return nil, fmt.Errorf("subnode not found for basePath %q", basePath)
	}

	return node, nil
}
