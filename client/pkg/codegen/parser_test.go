package codegen

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name          string
		source        []byte
		structName    string
		startLine     int
		expected      *Struct
		expectedError error
	}{
		{
			name: "happy path parsing",
			source: []byte(`
package testpackage

type DnsRecord struct {
	ID	 			   string` + " `codegen:\"id\"`" + `
	Name 			   string` + " `codegen:\"name\"`" + `
	GeneratedNumber	   string` + " `codegen:\"internal_id\"`" + `
	Enabled 		   bool` + " `codegen:\"enabled\"`" + `
	Omitted			   bool` + " `codegen:\"-\"`" + `
	ExplicitlyOmitted  bool` + " `codegen:\"-,omit\"`" + `
}
			`),

			expected: &Struct{
				Name: "DnsRecord",
				Fields: []*Field{
					{
						Name:       "ID",
						NameTarget: "id",
						Type:       NewType(TypeString),
					},
					{
						Name:       "Name",
						NameTarget: "name",
						Type:       NewType(TypeString),
					},
					{
						Name:       "GeneratedNumber",
						NameTarget: "internal_id",
						Type:       NewType(TypeString),
					},
					{
						Name:       "Enabled",
						NameTarget: "enabled",
						Type:       NewType(TypeBool),
					},
				},
			},
		},
		{
			name: "all tags are passed, not just known",
			source: []byte(`
package testpackage

type DnsRecord struct {
	ID	 			   string` + " `codegen:\"id\"`" + `
	Name 			   string` + " `codegen:\"name\"`" + `
	GeneratedNumber	   string` + " `codegen:\"internal_id\"`" + `
	Enabled 		   bool` + " `codegen:\"enabled\"`" + `
	Omitted			   bool` + " `codegen:\"-\"`" + `
	ExplicitlyOmitted  bool` + " `codegen:\"-,omit\"`" + `
}
			`),

			expected: &Struct{
				Name: "DnsRecord",
				Fields: []*Field{
					{
						Name:       "ID",
						NameTarget: "id",
						Type:       NewType(TypeString),
					},
					{
						Name:       "Name",
						NameTarget: "name",
						Type:       NewType(TypeString),
					},
					{
						Name:       "GeneratedNumber",
						NameTarget: "internal_id",
						Type:       NewType(TypeString),
					},
					{
						Name:       "Enabled",
						NameTarget: "enabled",
						Type:       NewType(TypeBool),
					},
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fSet := token.NewFileSet()
			node, err := parser.ParseFile(fSet, "", tc.source, parser.ParseComments)
			if err != nil {
				t.Error(err)
			}

			result, err := parse(fSet, node, tc.startLine, tc.structName)
			// todo(maksym): this condition does not check the error type since we don't have specific errors yet
			if (tc.expectedError == nil) != (err == nil) {
				t.Errorf("expected error to be %v, got %v", tc.expectedError, err)
			}
			if err != nil {
				return
			}
			assert.Equal(t, tc.expected, result)
		})
	}
}
