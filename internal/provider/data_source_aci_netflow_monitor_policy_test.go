// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNetflowMonitorPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigNetflowMonitorPolDataSourceDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_netflow_monitor_policy.test", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("data.aci_netflow_monitor_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_netflow_monitor_policy.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_monitor_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_monitor_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_monitor_policy.test", "owner_tag", ""),
				),
			},
			{
				Config:      testConfigNetflowMonitorPolNotExistingFvTenant,
				ExpectError: regexp.MustCompile("Failed to read aci_netflow_monitor_policy data source"),
			},
		},
	})
}

const testConfigNetflowMonitorPolDataSourceDependencyWithFvTenant = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
data "aci_netflow_monitor_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
  depends_on = [aci_netflow_monitor_policy.test]
}
`

const testConfigNetflowMonitorPolNotExistingFvTenant = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
data "aci_netflow_monitor_policy" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor_non_existing"
  depends_on = [aci_netflow_monitor_policy.test]
}
`
