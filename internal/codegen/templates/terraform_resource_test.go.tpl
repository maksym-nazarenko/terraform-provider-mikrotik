package mikrotik

import (
	{{ range $import := .Imports -}}
		"{{ $import }}"
	{{ end }}
)

{{ $resourceNameTitle := .ClientResource}}
{{ $resourceNameIdent := .ClientResource | snakeCase }}
{{ $resourceType := $resourceNameTitle | snakeCase | printf "mikrotik_%s" }}

func TestAcc{{$resourceNameTitle}}_basic(t *testing.T) {
	resourceName := "{{$resourceType}}.testacc_{{$resourceNameIdent}}"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProtoV5ProviderFactories,
		CheckDestroy:             testAccCheck{{$resourceNameTitle}}Destroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "{{$resourceType}}" "testacc_{{$resourceNameIdent}}" {
						{{ range $x := .Fields -}}
						{{if and $x.Computed (not $x.Optional) -}}{{continue}}{{end -}}
							{{ $x.AttributeName }} = {{$x.Type.Name | sampleData}}
						{{ end }}
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAcc{{$resourceNameTitle}}Exists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName: resourceName,
				ImportState:  true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceName)
					}
					return rs.Primary.Attributes["{{.IDField.AttributeName}}"], nil
				},
			},
		},
	})
}

func testAccCheck{{$resourceNameTitle}}Destroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "{{ $resourceType }}" {
			continue
		}

		remoteRecord, err := c.Find{{$resourceNameTitle}}(rs.Primary.Attributes["{{.IDField.AttributeName}}"])

		if !client.IsNotFoundError(err) && err != nil {
			return err
		}

		if remoteRecord != nil {
			return fmt.Errorf("remote record (%s) still exists", remoteRecord.ID())
		}

	}
	return nil
}

func testAcc{{$resourceNameTitle}}Exists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("%s does not exist in the statefile", resourceName)
		}

		c := client.NewClient(client.GetConfigFromEnv())
		record, err := c.Find{{$resourceNameTitle}}(rs.Primary.Attributes["{{.IDField.AttributeName}}"])
		if err != nil {
			return fmt.Errorf("Unable to get remote record for %s: %v", resourceName, err)
		}

		if record == nil {
			return fmt.Errorf("Unable to get the remote record %s", resourceName)
		}

		return nil
	}
}
