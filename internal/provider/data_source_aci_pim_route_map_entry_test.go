// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourcePimRouteMapEntryWithPimRouteMapPol(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigPimRouteMapEntryDataSourceDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "order", "1"),
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "action", "permit"),
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "name", ""),
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.aci_pim_route_map_entry.test", "source_ip", "0.0.0.0"),
				),
			},
			{
				Config:      testConfigPimRouteMapEntryNotExistingPimRouteMapPol,
				ExpectError: regexp.MustCompile("Failed to read aci_pim_route_map_entry data source"),
			},
		},
	})
}

const testConfigPimRouteMapEntryDataSourceDependencyWithPimRouteMapPol = testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPol + `
data "aci_pim_route_map_entry" "test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
  depends_on = [aci_pim_route_map_entry.test]
}
`

const testConfigPimRouteMapEntryNotExistingPimRouteMapPol = testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPol + `
data "aci_pim_route_map_entry" "test_non_existing" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "2"
  depends_on = [aci_pim_route_map_entry.test]
}
`
