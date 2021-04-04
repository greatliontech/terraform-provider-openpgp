package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceKey(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceKey,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"openpgp_key.foo", "name", regexp.MustCompile("^John Doe$")),
					resource.TestMatchResourceAttr(
						"openpgp_key.foo", "email", regexp.MustCompile("^john.doe@example.com$")),
					resource.TestMatchResourceAttr(
						"openpgp_key.foo", "private_key_armor", regexp.MustCompile("PGP PRIVATE KEY BLOCK")),
					resource.TestMatchResourceAttr(
						"openpgp_key.foo", "public_key_armor", regexp.MustCompile("PGP PUBLIC KEY BLOCK")),
				),
			},
		},
	})
}

const testAccResourceKey = `
resource "openpgp_key" "foo" {
  name = "John Doe"
  email = "john.doe@example.com"
}
`
