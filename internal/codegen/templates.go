package codegen

import _ "embed"

const (
	generatedNotice = `// This code was generated. Review it carefully.

`
)

var (
	//go:embed templates/terraform_resource.go.tpl
	resourceTemplate string

	terraformResourceDefinitionTemplate string = generatedNotice + resourceTemplate
)
