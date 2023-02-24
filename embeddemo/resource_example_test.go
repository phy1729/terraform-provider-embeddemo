package embeddemo_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExample_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "embeddemo_example" "test" {
						bar = "bar"
						foo = "foo"
						id  = "id"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("embeddemo_example.test", "bar", "bar"),
					resource.TestCheckResourceAttr("embeddemo_example.test", "foo", "foo"),
					resource.TestCheckResourceAttr("embeddemo_example.test", "id", "id"),
				),
			},
		},
	})
}
