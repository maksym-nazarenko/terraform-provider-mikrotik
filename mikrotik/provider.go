package mikrotik

import (
	"github.com/go-routeros/routeros"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MIKROTIK_HOST", nil),
				Description: "Hostname of the mikrotik router",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MIKROTIK_USER", nil),
				Description: "User account for mikrotik api",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MIKROTIK_PASSWORD", nil),
				Description: "Password for mikrotik api",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mikrotik_dns_record": resourceRecord(),
		},
		// TODO: do i need a configure func?
		ConfigureFunc: mikrotikConfigure,
	}
}

func mikrotikConfigure(d *schema.ResourceData) (c interface{}, err error) {
	address := d.Get("host").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	c = mikrotikConn{
		host:     address,
		username: username,
		password: password,
	}
	return
}

type mikrotikConn struct {
	host     string
	username string
	password string
}

func getMikrotikClient(m interface{}) (c *routeros.Client, err error) {
	conn := m.(mikrotikConn)
	address := conn.host
	username := conn.username
	password := conn.password
	c, err = routeros.Dial(address, username, password)
	return
}