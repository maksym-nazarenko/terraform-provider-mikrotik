package codegen

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
	"text/template"

	"github.com/ddelnano/terraform-provider-mikrotik/client/internal/codegen/utils"
	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/inspect"
)

type (
	// ArgumentDesciptor is the interface the template logic interacts with.
	ArgumentDesciptor interface {
		Name() string
		Options() []string
		Type() string
	}

	PropertyDescriptor interface {
		Name() string
		Type() string
		Readonly() bool
	}

	argumentWrapper struct {
		arg          *inspect.Argument
		possibleType string
	}

	propertyWrapper struct {
		property *inspect.Property
	}

	definitionData struct {
		CommandBasePath    string
		ResourceName       string
		Arguments          []ArgumentDesciptor
		ReadonlyProperties []PropertyDescriptor
	}
)

func GenerateMikrotikResource(resourceName string, node *inspect.Node, w io.Writer) error {
	if node == nil {
		return errors.New("node object cannot be nil")
	}

	addCmd, ok := node.ChildrenMap["add"]
	if !ok {
		return fmt.Errorf("command 'add' was not found in definition of resource %q", node.Self)
	}

	arguments := make([]ArgumentDesciptor, 0, len(addCmd.Arguments))
	for i := range addCmd.Arguments {
		arguments = append(arguments, &argumentWrapper{arg: addCmd.Arguments[i]})
	}

	readonlyProperties := make([]PropertyDescriptor, 0, len(node.ReadonlyPropertiesMap))
	for name := range node.ReadonlyPropertiesMap {
		readonlyProperties = append(readonlyProperties, &propertyWrapper{property: node.ReadonlyPropertiesMap[name]})
	}

	data := &definitionData{
		CommandBasePath:    node.Self,
		ResourceName:       resourceName,
		Arguments:          arguments,
		ReadonlyProperties: readonlyProperties,
	}
	return generateCode(
		w,
		"resource",
		mikrotikResourceDefinitionTemplate,
		data,
	)
}

func generateCode(w io.Writer, templateName, templateBody string, templateData any) error {
	t := template.New(templateName)
	t.Funcs(template.FuncMap{
		"pascalCase": utils.PascalCase,
	})

	if _, err := t.Parse(templateBody); err != nil {
		return err
	}

	if err := t.Execute(w, templateData); err != nil {
		return err
	}

	return nil
}

func (aw *argumentWrapper) Name() string {
	return aw.arg.Name
}

func (aw *argumentWrapper) Options() []string {
	return aw.arg.Options
}

func (aw *argumentWrapper) Type() string {
	if aw.possibleType == "" {
		aw.possibleType = aw.guessType()
	}

	return aw.possibleType
}

func (aw *argumentWrapper) guessType() string {
	if len(aw.arg.Options) == 2 && slices.Contains(aw.arg.Options, "no") && slices.Contains(aw.arg.Options, "yes") {
		return "bool"
	}

	if strings.ToLower(aw.arg.Name) == "ttl" {
		return "types.MikrotikDuration"
	}

	if len(aw.arg.Options) == 1 && strings.ToLower(aw.arg.Options[0]) == "none" {
		// return "types.MikrotikNoneAware"
		return "string"
	}

	return "string"
}

func (pw *propertyWrapper) Name() string {
	return pw.property.Name
}

func (pw *propertyWrapper) Type() string {
	return "string"
}

func (pw *propertyWrapper) Readonly() bool {
	return true
}
