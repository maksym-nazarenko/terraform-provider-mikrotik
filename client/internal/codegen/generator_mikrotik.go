package codegen

import (
	"errors"
	"fmt"
	"io"
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

func GenerateMikrotikResourceTest(s *Struct, w io.Writer) error {
	data := struct {
		ResourceName string
		Fields       []*Field
	}{
		ResourceName: s.Name,
		Fields:       s.Fields,
	}

	return generateCode(
		w,
		"resource-test",
		mikrotikResourceTestDefinitionTemplate,
		data,
	)
}

func generateCode(w io.Writer, templateName, templateBody string, templateData any) error {
	t := template.New(templateName)
	t.Funcs(template.FuncMap{
		"pascalCase": utils.PascalCase,
		"snakeCase":  utils.ToSnakeCase,
		"sampleData": sampleData,
	})

	if _, err := t.Parse(templateBody); err != nil {
		return err
	}

	if err := t.Execute(w, templateData); err != nil {
		return err
	}

	return nil
}

// sampleData generates sample value for provided type.
func sampleData(typeName string) string {
	switch typeName {
	case string(TypeString):
		return `"sample"`
	case string(TypeInt64):
		return "42"
	case string(TypeBool):
		return "false"
	default:
		return `"` + string(TypeUnknown) + `"`
	}
}
