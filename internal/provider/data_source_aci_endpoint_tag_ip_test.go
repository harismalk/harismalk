// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvEpIpTagWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "5.2(1g)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvEpIpTagDataSourceDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_endpoint_tag_ip.test", "ip", "10.0.0.2"),
					resource.TestCheckResourceAttr("data.aci_endpoint_tag_ip.test", "vrf_name", "test_ctx_name"),
					resource.TestCheckResourceAttr("data.aci_endpoint_tag_ip.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_endpoint_tag_ip.test", "id_attribute", "0"),
					resource.TestCheckResourceAttr("data.aci_endpoint_tag_ip.test", "name", ""),
					resource.TestCheckResourceAttr("data.aci_endpoint_tag_ip.test", "name_alias", ""),
				),
			},
			{
				Config:      testConfigFvEpIpTagNotExistingFvTenant,
				ExpectError: regexp.MustCompile("Failed to read aci_endpoint_tag_ip data source"),
			},
		},
	})
}

const testConfigFvEpIpTagDataSourceDependencyWithFvTenant = testConfigFvEpIpTagMinDependencyWithFvTenant + `
data "aci_endpoint_tag_ip" "test" {
  parent_dn = aci_tenant.test.id
  ip = "10.0.0.2"
  vrf_name = "test_ctx_name"
  depends_on = [aci_endpoint_tag_ip.test]
}
`

const testConfigFvEpIpTagNotExistingFvTenant = testConfigFvEpIpTagMinDependencyWithFvTenant + `
data "aci_endpoint_tag_ip" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  ip = "10.0.1.2"
  vrf_name = "non_existing_ctx_name"
  depends_on = [aci_endpoint_tag_ip.test]
}
`
