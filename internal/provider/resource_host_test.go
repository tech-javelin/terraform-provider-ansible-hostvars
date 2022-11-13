package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceHost(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceHost,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"ansible_role_bind_host.test", "domain", regexp.MustCompile("^foo.tld")),
				),
			},
		},
	})
}

const testAccResourceHost = `
resource "ansible_role_bind_host" "test" {
  domain = "foo.tld"
  hostname = "test"
  ip = "1.2.3.4"
  aliases = []
}
`
