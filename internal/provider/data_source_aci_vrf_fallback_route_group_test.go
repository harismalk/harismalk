// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvFBRGroupWithFvCtx(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvFBRGroupDataSourceDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("data.aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_vrf_fallback_route_group.test", "name_alias", ""),
				),
			},
			{
				Config:      testConfigFvFBRGroupNotExistingFvCtx,
				ExpectError: regexp.MustCompile("Failed to read aci_vrf_fallback_route_group data source"),
			},
		},
	})
}

const testConfigFvFBRGroupDataSourceDependencyWithFvCtx = testConfigFvFBRGroupMinDependencyWithFvCtx + `
data "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
  depends_on = [aci_vrf_fallback_route_group.test]
}
`

const testConfigFvFBRGroupNotExistingFvCtx = testConfigFvFBRGroupMinDependencyWithFvCtx + `
data "aci_vrf_fallback_route_group" "test_non_existing" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group_non_existing"
  depends_on = [aci_vrf_fallback_route_group.test]
}
`
