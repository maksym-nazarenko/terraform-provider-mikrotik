package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"

	codegenClient "github.com/ddelnano/terraform-provider-mikrotik/client/pkg/codegen"
	"github.com/ddelnano/terraform-provider-mikrotik/internal/codegen"
)

func subcommandResource(args []string) error {
	var (
		sourceFile string
		structName string
		outFile    string
	)

	fs := flag.NewFlagSet("terraform", flag.ExitOnError)

	fs.StringVar(&sourceFile, "srcFile", "", "`path` to source file to parse struct from")
	fs.StringVar(&outFile, "outFile", "-", "`path` to write generated code to. Use '-' for stdout")
	fs.StringVar(&structName, "structName", "", "name of a struct to process")
	if err := fs.Parse(args); err != nil {
		return err
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

	s, err := codegenClient.ParseFile(sourceFile, startLine, structName)
	if err != nil {
		return err
	}

	// If struct name is not provider, use one found in the parsed file.
	if structName == "" {
		structName = s.Name
	}

	type (
		generatorPostHookFunc func([]byte) ([]byte, error)
	)
	var (
		buf    bytes.Buffer
		result []byte
	)

	if err := codegen.GenerateResource(s, &buf); err != nil {
		return err
	}

	writeHooks := []generatorPostHookFunc{codegenClient.SourceFormatHook}

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

func usage(appName string) {
	fmt.Fprintf(os.Stdout,
		`Usage: %s <type> [flags]

<type> is one of:
  resource	- generate a Terraform resource
  test		- generate a test-file for an existing Terraform resource

run %s <type> -h for more information about a specific type flags`,
		appName, appName,
	)
}
