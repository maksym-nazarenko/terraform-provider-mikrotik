package codegen

const (
	generatedNotice = "// This code was generated. Review it carefully."

	terraformResourceDefinitionTemplate = generatedNotice + `

package mikrotik

import (
	{{ range $import := .Imports -}}
		"{{ $import }}"
	{{ end }}
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"

)
{{ $resourceStructName := .ClientResource | firstLower}}
{{ $clientResourceName := .ClientResource }}
type {{$resourceStructName}} struct {
	client *client.Mikrotik
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &{{$resourceStructName}}{}
	_ resource.ResourceWithConfigure   = &{{$resourceStructName}}{}
	_ resource.ResourceWithImportState = &{{$resourceStructName}}{}
)

// New{{$clientResourceName}}Resource is a helper function to simplify the provider implementation.
func New{{$clientResourceName}}Resource() resource.Resource {
	return &{{$resourceStructName}}{}
}

func (r *{{$resourceStructName}}) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Mikrotik)
}

// Metadata returns the resource type name.
func (r *{{$resourceStructName}}) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{$clientResourceName| snakeCase}}"
}

// Schema defines the schema for the resource.
func (s *{{$resourceStructName}}) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Creates a MikroTik {{$clientResourceName}}.",
		Attributes: map[string]schema.Attribute{
			{{range .Fields -}}
			"{{.AttributeName}}": schema.{{.Type.Name}}Attribute{
				Required: {{.Required}},
				Optional: {{.Optional}},
				Computed: {{.Computed}},
				{{if .Computed -}}
				PlanModifiers: []planmodifier.{{.Type.Name}}{
					{{.Type.Name | lowerCase}}planmodifier.UseStateForUnknown(),
				},
				{{- end}}
				Description: "",
			},
			{{end}}
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *{{$resourceStructName}}) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var terraformModel {{$resourceStructName}}Model
	var mikrotikModel client.{{$clientResourceName}}
	GenericCreateResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

// Read refreshes the Terraform state with the latest data.
func (r *{{$resourceStructName}}) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var terraformModel {{$resourceStructName}}Model
	var mikrotikModel client.{{$clientResourceName}}
	GenericReadResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *{{$resourceStructName}}) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var terraformModel {{$resourceStructName}}Model
	var mikrotikModel client.{{$clientResourceName}}
	GenericUpdateResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *{{$resourceStructName}}) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var terraformModel {{$resourceStructName}}Model
	var mikrotikModel client.{{$clientResourceName}}
	GenericDeleteResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

func (r *{{$resourceStructName}}) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("{{.IDField.AttributeName}}"), req, resp)
}

type {{$resourceStructName}}Model struct {
	{{range .Fields -}}
	{{.Name}}        tftypes.{{.Type.Name}} ` + "`" + `tfsdk:"{{.AttributeName}}"` + "`" + `
	{{end}}
}
`
)
