package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

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
	appName := path.Base(os.Args[0])
	if len(os.Args) < 2 {
		usage(appName)
		return nil
	}
	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "resource":
		return subcommandResource(args)
	case "-h", "--help":
		usage(appName)
		return nil
	default:
		usage(appName)
		return fmt.Errorf("unknown subcommand: %s", subcommand)
	}
}

func subcommandResource(args []string) error {
	var (
		queryRemote    bool
		commandBase    string
		resourceName   string
		definitionFile string
	)

	fs := flag.NewFlagSet("resource", flag.ExitOnError)
	fs.BoolVar(&queryRemote, "query", false, "query remote system to fetch resource definition")
	fs.StringVar(&commandBase, "basePath", "", "resource base path to generate code for")
	fs.StringVar(&definitionFile, "definitionFile", "", "`path` to definition file in JSON format")
	fs.StringVar(&resourceName, "resourceName", "", "use this name in code, otherwise it is generated using basePath")
	if err := fs.Parse(args); err != nil {
		return err
	}

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
			return fmt.Errorf("could not parse resource definition file: %w", err)
		}
	}

	if rootNode == nil {
		return fmt.Errorf("could not build root node")
	}

	if resourceName == "" {
		resourceName = utils.PascalCase(rootNode.Self)
	}

	type (
		generatorPostHookFunc func([]byte) ([]byte, error)
	)
	var (
		buf    bytes.Buffer
		result []byte
		err    error
	)

	if err := codegen.GenerateMikrotikResource(resourceName, rootNode, &buf); err != nil {
		return err
	}

	writeHooks := []generatorPostHookFunc{codegen.SourceFormatHook}

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

func usage(appName string) {
	fmt.Fprintf(os.Stdout,
		`Usage: %s <type> [flags]

<type> is one of:
  resource  - generate a MikroTik resource to be used by the API client

run %s <type> -h for more information about a specific type flags`,
		appName, appName,
	)
}
