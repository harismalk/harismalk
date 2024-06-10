// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceNetflowRecordPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigNetflowRecordPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name", "netfow_record"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "collect_paramaters.0", "src-intf"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "match_parameters.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_tag", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigNetflowRecordPolAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name", "netfow_record"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "collect_paramaters.0", "count-bytes"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "collect_paramaters.1", "src-intf"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "match_parameters.0", "dst-ip"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "match_parameters.1", "src-ip"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_tag", "owner_tag"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigNetflowRecordPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name", "netfow_record"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigNetflowRecordPolResetDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name", "netfow_record"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "collect_paramaters.0", "src-intf"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "match_parameters.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_tag", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_netflow_record_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name", "netfow_record"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "collect_paramaters.0", "src-intf"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "match_parameters.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_tag", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigNetflowRecordPolChildrenDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name", "netfow_record"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "collect_paramaters.0", "src-intf"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "match_parameters.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_netflow_record_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name", "netfow_record"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "collect_paramaters.0", "src-intf"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "match_parameters.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigNetflowRecordPolChildrenRemoveFromConfigDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigNetflowRecordPolChildrenRemoveOneDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigNetflowRecordPolChildrenRemoveAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_record_policy.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigNetflowRecordPolMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_record_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record"
}
`

const testConfigNetflowRecordPolAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_record_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record"
  annotation = "annotation"
  collect_paramaters = ["count-bytes", "src-intf"]
  description = "description"
  match_parameters = ["dst-ip", "src-ip"]
  name_alias = "name_alias"
  owner_key = "owner_key"
  owner_tag = "owner_tag"
}
`

const testConfigNetflowRecordPolResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_record_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record"
  annotation = "orchestrator:terraform"
  collect_paramaters = ["src-intf"]
  description = ""
  match_parameters = []
  name_alias = ""
  owner_key = ""
  owner_tag = ""
}
`
const testConfigNetflowRecordPolChildrenDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_record_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record"
  annotations = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  tags = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
}
`

const testConfigNetflowRecordPolChildrenRemoveFromConfigDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_record_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record"
}
`

const testConfigNetflowRecordPolChildrenRemoveOneDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_record_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record"
  annotations = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
}
`

const testConfigNetflowRecordPolChildrenRemoveAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_record_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_record"
  annotations = []
  tags = []
}
`
