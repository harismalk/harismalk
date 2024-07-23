// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNetflowRecordPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigNetflowRecordPolDataSourceDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "name", "netfow_record"),
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "collect_parameters.#", "1"),
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "collect_parameters.0", "src-intf"),
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "match_parameters.#", "0"),
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_record_policy.test", "owner_tag", ""),
				),
			},
			{
				Config:      testConfigNetflowRecordPolNotExistingFvTenant,
				ExpectError: regexp.MustCompile("Failed to read aci_netflow_record_policy data source"),
			},
		},
	})
}

const testConfigNetflowRecordPolDataSourceDependencyWithFvTenant = testConfigNetflowRecordPolMinDependencyWithFvTenant + `
data "aci_netflow_record_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record"
  depends_on = [aci_netflow_record_policy.test]
}
`

const testConfigNetflowRecordPolNotExistingFvTenant = testConfigFvTenantMin + `
data "aci_netflow_record_policy" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record_non_existing"
}
`
