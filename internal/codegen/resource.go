package codegen

import (
	"html/template"
	"io"
	"strings"

	clientCodegen "github.com/ddelnano/terraform-provider-mikrotik/client/pkg/codegen"
	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/codegen/utils"
)

var (
	terraformResourceImports = []string{
		"context",
		"github.com/ddelnano/terraform-provider-mikrotik/client",
		"github.com/hashicorp/terraform-plugin-framework/path",
		"github.com/hashicorp/terraform-plugin-framework/resource",
		"github.com/hashicorp/terraform-plugin-framework/resource/schema",
		"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier",
		"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier",
	}
)

type (
	FieldType string

	TypeDescriptor interface {
		Name() string
	}

	Field struct {
		Name          string
		AttributeName string
		Type          TypeDescriptor
		Required      bool
		Optional      bool
		Computed      bool
	}

	templateData struct {
		ClientResource string
		Imports        []string
		Fields         []*Field
		IDField        *Field
	}
)

// GenerateResource generates Terraform resource and writes it to specified output
func GenerateResource(s *clientCodegen.Struct, w io.Writer) error {
	fields, err := convertClientFields(s.Fields)
	if err != nil {
		return err
	}

	data := &templateData{
		ClientResource: s.Name,
		Imports:        terraformResourceImports,
		IDField: &Field{
			Name:          "ID",
			Type:          FieldType("String"),
			AttributeName: "id",
		},
		Fields: fields,
	}

	return generateCode(w,
		"resource",
		terraformResourceDefinitionTemplate,
		data,
	)
}

func generateCode(w io.Writer, templateName, templateBody string, templateData any) error {
	t := template.New(templateName)
	t.Funcs(template.FuncMap{
		"pascalCase": utils.PascalCase,
		"snakeCase":  utils.ToSnakeCase,
		"firstLower": utils.FirstLower,
		"lowerCase":  strings.ToLower,
	})

	if _, err := t.Parse(templateBody); err != nil {
		return err
	}

	if err := t.Execute(w, templateData); err != nil {
		return err
	}

	return nil
}

func convertClientFields(clientFields []*clientCodegen.Field) ([]*Field, error) {
	result := make([]*Field, 0)

	for _, v := range clientFields {
		result = append(result, &Field{
			Name:          v.Name,
			AttributeName: v.NameTarget,
			Type:          FieldType(structTypeToTerraformType(v.Type.Type())),
			Computed:      v.Readonly,
		})
	}

	return result, nil
}

func (ft FieldType) Name() string {
	return string(ft)
}

func structTypeToTerraformType(typ string) string {
	switch typ {
	case "string":
		return "String"
	case "int64":
		return "Int64"
	case "bool":
		return "Bool"
	default:
		return "unknown"
	}
}
