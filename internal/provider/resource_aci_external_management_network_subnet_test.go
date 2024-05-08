// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceMgmtSubnetWithMgmtInstP(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigMgmtSubnetMinDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "ip", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "description", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name_alias", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigMgmtSubnetAllDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "ip", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name", "name"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name_alias", "name_alias"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigMgmtSubnetMinDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "ip", "1.1.1.0/24"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigMgmtSubnetResetDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "ip", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "description", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name_alias", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_external_management_network_subnet.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "ip", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "description", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name_alias", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigMgmtSubnetChildrenDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "ip", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "description", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_external_management_network_subnet.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "ip", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "description", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigMgmtSubnetChildrenRemoveFromConfigDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigMgmtSubnetChildrenRemoveOneDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigMgmtSubnetChildrenRemoveAllDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_external_management_network_subnet.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigMgmtSubnetMinDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_external_management_network_subnet" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  ip = "1.1.1.0/24"
}
`

const testConfigMgmtSubnetAllDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_external_management_network_subnet" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  ip = "1.1.1.0/24"
  annotation = "annotation"
  description = "description"
  name = "name"
  name_alias = "name_alias"
}
`

const testConfigMgmtSubnetResetDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_external_management_network_subnet" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  ip = "1.1.1.0/24"
  annotation = "orchestrator:terraform"
  description = ""
  name = ""
  name_alias = ""
}
`
const testConfigMgmtSubnetChildrenDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_external_management_network_subnet" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  ip = "1.1.1.0/24"
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

const testConfigMgmtSubnetChildrenRemoveFromConfigDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_external_management_network_subnet" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  ip = "1.1.1.0/24"
}
`

const testConfigMgmtSubnetChildrenRemoveOneDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_external_management_network_subnet" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  ip = "1.1.1.0/24"
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

const testConfigMgmtSubnetChildrenRemoveAllDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_external_management_network_subnet" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  ip = "1.1.1.0/24"
  annotations = []
  tags = []
}
`
