// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvEpMacTagWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvEpMacTagMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "bd_name", "test_bd_name"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "mac", "00:00:00:00:00:01"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "id_attribute", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name", ""),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name_alias", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvEpMacTagAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "bd_name", "test_bd_name"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "mac", "00:00:00:00:00:01"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "id_attribute", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name", "name"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name_alias", "name_alias"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvEpMacTagMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "bd_name", "test_bd_name"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "mac", "00:00:00:00:00:01"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvEpMacTagResetDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "bd_name", "test_bd_name"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "mac", "00:00:00:00:00:01"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "id_attribute", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name", ""),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name_alias", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_endpoint_tag_mac.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "bd_name", "test_bd_name"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "mac", "00:00:00:00:00:01"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "id_attribute", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name", ""),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name_alias", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigFvEpMacTagChildrenDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "bd_name", "test_bd_name"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "mac", "00:00:00:00:00:01"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "id_attribute", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name", ""),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_endpoint_tag_mac.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "bd_name", "test_bd_name"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "mac", "00:00:00:00:00:01"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "id_attribute", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name", ""),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvEpMacTagChildrenRemoveFromConfigDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvEpMacTagChildrenRemoveOneDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.deletable_child", "true"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.deletable_child", "true"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvEpMacTagChildrenRemoveAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_tag_mac.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvEpMacTagMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_endpoint_tag_mac" "test" {
  parent_dn = aci_tenant.test.id
  bd_name = "test_bd_name"
  mac = "00:00:00:00:00:01"
}
`

const testConfigFvEpMacTagAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_endpoint_tag_mac" "test" {
  parent_dn = aci_tenant.test.id
  bd_name = "test_bd_name"
  mac = "00:00:00:00:00:01"
  annotation = "annotation"
  id_attribute = "1"
  name = "name"
  name_alias = "name_alias"
}
`

const testConfigFvEpMacTagResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_endpoint_tag_mac" "test" {
  parent_dn = aci_tenant.test.id
  bd_name = "test_bd_name"
  mac = "00:00:00:00:00:01"
  annotation = "orchestrator:terraform"
  id_attribute = "0"
  name = ""
  name_alias = ""
}
`
const testConfigFvEpMacTagChildrenDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_endpoint_tag_mac" "test" {
  parent_dn = aci_tenant.test.id
  bd_name = "test_bd_name"
  mac = "00:00:00:00:00:01"
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

const testConfigFvEpMacTagChildrenRemoveFromConfigDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_endpoint_tag_mac" "test" {
  parent_dn = aci_tenant.test.id
  bd_name = "test_bd_name"
  mac = "00:00:00:00:00:01"
}
`

const testConfigFvEpMacTagChildrenRemoveOneDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_endpoint_tag_mac" "test" {
  parent_dn = aci_tenant.test.id
  bd_name = "test_bd_name"
  mac = "00:00:00:00:00:01"
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

const testConfigFvEpMacTagChildrenRemoveAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_endpoint_tag_mac" "test" {
  parent_dn = aci_tenant.test.id
  bd_name = "test_bd_name"
  mac = "00:00:00:00:00:01"
  annotations = [
  ]
  tags = [
  ]
}
`
