package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/ddelnano/terraform-provider-mikrotik/client/internal/codegen"
	codegenPkg "github.com/ddelnano/terraform-provider-mikrotik/client/pkg/codegen"
	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/codegen/utils"
	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/inspect"
)

func subcommandResource(args []string) error {
	var (
		queryRemote    bool
		commandBase    string
		resourceName   string
		definitionFile string
		outFile        string
	)

	fs := flag.NewFlagSet("resource", flag.ExitOnError)
	fs.BoolVar(&queryRemote, "query", false, "query remote system to fetch resource definition")
	fs.StringVar(&commandBase, "basePath", "", "resource base path to generate code for")
	fs.StringVar(&definitionFile, "definitionFile", "", "`path` to definition file in JSON format")
	fs.StringVar(&resourceName, "resourceName", "", "use this name in code, otherwise it is generated using basePath")
	fs.StringVar(&outFile, "outFile", "-", "`path` to write generated code to. Use '-' for stdout")
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
		var rootNodeDefinition *inspect.Node
		var err error
		if rootNodeDefinition, err = readDefinitionFile(definitionFile); err != nil {
			return err
		}
		rootNode, err = findSubNode(rootNodeDefinition, commandBase)
		if err != nil {
			return err
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

	writeHooks := []generatorPostHookFunc{codegenPkg.SourceFormatHook}

	result = buf.Bytes()
	for _, h := range writeHooks {
		result, err = h(result)
		if err != nil {
			return err
		}
	}

	writeOut := func(data []byte) error {
		_, err := fmt.Fprint(os.Stdout, string(data))
		return err
	}

	if outFile != "-" {
		writeOut = func(data []byte) error {
			return os.WriteFile(outFile, data, 0644)
		}
	}

	return writeOut(result)
}

func subcommandTest(args []string) error {
	var (
		sourceFile string
		structName string
		outputFile string
	)

	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.StringVar(&sourceFile, "sourceFile", "", "`path` to source file to generate test for")
	fs.StringVar(&structName, "structName", "", "name of the struct to search in the source file, if not provided the first struct found will be used")
	fs.StringVar(&outputFile, "outFile", "-", "`path` to output file for the generated test. Use '-' for stdout")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if sourceFile == "" {
		return errors.New("no source file provided via -sourceFile flag")
	}

	startLine := 1
	lineStr := os.Getenv("GOLINE")
	if lineStr != "" {
		lineInt, err := strconv.Atoi(lineStr)
		if err != nil {
			return fmt.Errorf("fail to parse GOLINE: %v", err.Error())
		}
		startLine = lineInt
	}

	s, err := codegenPkg.ParseFile(sourceFile, startLine, structName)
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	if err := codegen.GenerateMikrotikResourceTest(s, &buf); err != nil {
		return err
	}

	writeOut := func(data []byte) error {
		_, err := fmt.Fprint(os.Stdout, string(data))
		return err
	}

	if outputFile != "-" {
		writeOut = func(data []byte) error {
			return os.WriteFile(outputFile, data, 0644)
		}
	}

	return writeOut(buf.Bytes())
}

func usage(appName string) {
	fmt.Fprintf(os.Stdout,
		`Usage: %s <type> [flags]

<type> is one of:
  resource	- generate a MikroTik resource to be used by the API client
  test		- generate a test-file for a MikroTik resource

run %s <type> -h for more information about a specific type flags`,
		appName, appName,
	)
}
