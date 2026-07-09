package codegen

const (
	generatedNotice = "// This code was auto-generated."

	mikrotikResourceDefinitionTemplate = generatedNotice + `

package client

import (
	"github.com/ddelnano/terraform-provider-mikrotik/client/types"
	"github.com/go-routeros/routeros/v3"
)
{{$resourceName := .ResourceName}}
{{ range $arg := .Arguments }}
{{- if and (gt ($arg.Options | len) 0) (eq $arg.Type "string") }}
// Constants for {{$arg.Name}}
const (
	{{- range $opt := .Options}}
	{{$resourceName}}_{{$arg.Name | pascalCase}}_{{$opt | pascalCase}} = "{{$opt}}"
	{{- end}}
)
{{end}}
{{- end}}

// {{.ResourceName}} defines resource
type {{.ResourceName}} struct {
	Id string ` + "`" + `mikrotik:".id" codegen:"id"` + "`" + `
	{{range $field := .Arguments -}}
		{{$field.Name | pascalCase}} {{$field.Type}} ` + "`" + `mikrotik:"{{$field.Name}}" codegen:"{{$field.Name | snakeCase}}"` + "`" + `
	{{end}}

	{{range $field := .ReadonlyProperties -}}
		{{$field.Name | pascalCase}} {{$field.Type}} ` + "`" + `mikrotik:"{{$field.Name}},readonly" codegen:"{{$field.Name | snakeCase}},readonly"` + "`" + `
	{{end}}
}

var _ Resource = (*{{.ResourceName}})(nil)

func (b *{{.ResourceName}}) ActionToCommand(a Action) string {
	return map[Action]string{
		Add:    "{{.CommandBasePath}}/add",
		Find:   "{{.CommandBasePath}}/print",
		Update: "{{.CommandBasePath}}/set",
		Delete: "{{.CommandBasePath}}/remove",
	}[a]
}

func (b *{{.ResourceName}}) IDField() string {
	return ".id"
}

func (b *{{.ResourceName}}) ID() string {
	return b.Id
}

func (b *{{.ResourceName}}) SetID(id string) {
	b.Id = id
}

// Uncomment extra methods to satisfy more interfaces

// Adder
func (b *{{.ResourceName}}) AfterAddHook(r *routeros.Reply) {
	b.Id = r.Done.Map["ret"]
}

// Finder
// func (b *{{.ResourceName}}) FindField() string {
// 	return "name"
// }

// func (b *{{.ResourceName}}) FindFieldValue() string {
// 	return b.Name
// }

// Deleter
// func (b *{{.ResourceName}}) DeleteField() string {
// 	return "numbers"
// }

// func (b *{{.ResourceName}}) DeleteFieldValue() string {
// 	return b.Id
// }


// Typed wrappers
func (c Mikrotik) Add{{.ResourceName}}(r *{{.ResourceName}}) (*{{.ResourceName}}, error) {
	res, err := c.Add(r)
	if err != nil {
		return nil, err
	}

	return res.(*{{.ResourceName}}), nil
}

func (c Mikrotik) Update{{.ResourceName}}(r *{{.ResourceName}}) (*{{.ResourceName}}, error) {
	res, err := c.Update(r)
	if err != nil {
		return nil, err
	}

	return res.(*{{.ResourceName}}), nil
}

func (c Mikrotik) Find{{.ResourceName}}(id string) (*{{.ResourceName}}, error) {
	res, err := c.Find(&{{.ResourceName}}{Id: id})
	if err != nil {
		return nil, err
	}

	return res.(*{{.ResourceName}}), nil
}

func (c Mikrotik) List{{.ResourceName}}() ([]{{.ResourceName}}, error) {
	res, err := c.List(&{{.ResourceName}}{})
	if err != nil {
		return nil, err
	}
	returnSlice := make([]{{.ResourceName}}, len(res))
	for i, v := range res {
		returnSlice[i] = *(v.(*{{.ResourceName}}))
	}

	return returnSlice, nil
}


func (c Mikrotik) Delete{{.ResourceName}}(id string) error {
	return c.Delete(&{{.ResourceName}}{Id: id})
}
`

	mikrotikResourceTestDefinitionTemplate = `
package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSmoke_{{.ResourceName}}(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	expectedResource := &{{.ResourceName}}{
		{{- range $field := .Fields }}
		{{- if eq $field.Name "Id" }}{{continue}}{{end}}
		// {{$field.Name}}: {{$field.Type.Type | sampleData}},
	{{- end }}
	}

	createdResource, err := c.Add{{.ResourceName}}(expectedResource)
	require.NoError(t, err)

	defer func(){
		err := c.Delete(createdResource)
		if !IsNotFoundError(err) {
			assert.NoError(t, err)
		}
	}()
	assert.NotEmpty(t, createdResource.Id)

	foundResource, err := c.Find(expectedResource)
	require.NoError(t, err)
	assert.Equal(t, createdResource, foundResource)

	// cleanup
	err = c.Delete(foundResource)
	assert.NoError(t, err)

	_, err = c.Find(expectedResource)
	assert.True(t, IsNotFoundError(err), "expected not found error, got: %v", err)
}
`
)
