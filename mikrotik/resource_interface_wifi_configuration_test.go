// This code was generated. Review it carefully.

package mikrotik

import (
	"fmt"
	"testing"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccInterfaceWiFiConfiguration_basic(t *testing.T) {
	resourceName := "mikrotik_interface_wi_fi_configuration.testacc_interface_wi_fi_configuration"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckInterfaceWiFiConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "mikrotik_interface_wi_fi_configuration" "testacc_interface_wi_fi_configuration" {
						comment = "sample"
						disabled = true
						name = "sample"
						ssid = "sample"
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccInterfaceWiFiConfigurationExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceName)
					}
					return rs.Primary.Attributes["id"], nil
				},
			},
		},
	})
}

func testAccCheckInterfaceWiFiConfigurationDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_interface_wi_fi_configuration" {
			continue
		}

		remoteRecord, err := c.FindInterfaceWiFiConfiguration(rs.Primary.Attributes["id"])

		if !client.IsNotFoundError(err) && err != nil {
			return err
		}

		if remoteRecord != nil {
			return fmt.Errorf("remote record (%s) still exists", remoteRecord.ID())
		}

	}
	return nil
}

func testAccInterfaceWiFiConfigurationExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("%s does not exist in the statefile", resourceName)
		}

		c := client.NewClient(client.GetConfigFromEnv())
		record, err := c.FindInterfaceWiFiConfiguration(rs.Primary.Attributes["id"])
		if err != nil {
			return fmt.Errorf("Unable to get remote record for %s: %v", resourceName, err)
		}

		if record == nil {
			return fmt.Errorf("Unable to get the remote record %s", resourceName)
		}

		return nil
	}
}
