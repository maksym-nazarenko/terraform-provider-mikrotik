package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/ddelnano/terraform-provider-mikrotik/client/internal/codegen"
	"github.com/ddelnano/terraform-provider-mikrotik/client/internal/codegen/utils"
	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/inspect"
)

func main() {
	if err := run(); err != nil {
		log.Println(fmt.Errorf("application failed: %w", err))
		os.Exit(1)
	}
}

func run() error {
	var (
		queryRemote    bool
		commandBase    string
		resourceName   string
		definitionFile string
	)

	flag.BoolVar(&queryRemote, "query", false, "query remote system to fetch resource definition")
	flag.StringVar(&commandBase, "basePath", "/", "resource base path to generate code for")
	flag.StringVar(&definitionFile, "definitionFile", "", "`path` to definition file in JSON format")
	flag.StringVar(&resourceName, "resourceName", "", "use this name in code, otherwise it is generated using basePath")
	flag.Parse()

	if len(commandBase) < 1 {
		return fmt.Errorf("commandBase cannot be empty")
	}
	if (len(definitionFile) > 0 && queryRemote) || (len(definitionFile) == 0 && !queryRemote) {
		return fmt.Errorf("at least (and at most) one of [query, definitionFile] flags must be set")
	}

	var rootNode *inspect.Node
	switch {

	case queryRemote:
		mc := client.NewClient(client.GetConfigFromEnv())
		c, err := mc.GetMikrotikClient()
		if err != nil {
			return err
		}

		rootNode, err = inspect.Do(c, &inspect.Config{
			Root: commandBase,
			// depth 2 is enough to get `/command/base/add` definition with all arguments
			Depth: 2,
		})
		if err != nil {
			return err
		}

	case definitionFile != "":
		fileBytes, err := os.ReadFile(definitionFile)
		if err != nil {
			return fmt.Errorf("could not read resource definition file: %w", err)
		}
		if err := json.Unmarshal(fileBytes, &rootNode); err != nil {
			return fmt.Errorf("coull not parse resource definition file: %w", err)
		}
	}

	if rootNode == nil {
		return fmt.Errorf("could not build root node")
	}

	var buf bytes.Buffer
	if resourceName == "" {
		resourceName = utils.PascalCase(rootNode.Self)
	}

	if err := codegen.GenerateMikrotikResource(resourceName, rootNode, &buf); err != nil {
		return err
	}

	type (
		generatorPostHookFunc func([]byte) ([]byte, error)
	)
	writeHooks := []generatorPostHookFunc{codegen.SourceFormatHook}

	var (
		result []byte
		err    error
	)

	result = buf.Bytes()
	for _, h := range writeHooks {
		result, err = h(result)
		if err != nil {
			return err
		}
	}

	out := os.Stdout
	if _, err := out.Write(result); err != nil {
		return err
	}

	return nil
}
