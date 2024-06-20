// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceMplsNodeSidPWithL3extLoopBackIfP(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigMplsNodeSidPAllDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", "node_sid_profile"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", "name_alias"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", "node_sid_profile"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigMplsNodeSidPResetDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", "node_sid_profile"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_l3out_node_sid_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", "node_sid_profile"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigMplsNodeSidPChildrenDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", "node_sid_profile"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_l3out_node_sid_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", "node_sid_profile"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigMplsNodeSidPChildrenRemoveFromConfigDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigMplsNodeSidPChildrenRemoveOneDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigMplsNodeSidPChildrenRemoveAllDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  loopback_address = "1.1.1.1"
  name = "node_sid_profile"
  segment_id = "1"
}
`

const testConfigMplsNodeSidPAllDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
  annotation = "annotation"
  description = "description"
  loopback_address = "1.1.1.1"
  name = "node_sid_profile"
  name_alias = "name_alias"
}
`

const testConfigMplsNodeSidPResetDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
  annotation = "orchestrator:terraform"
  description = ""
  loopback_address = "0.0.0.0"
  name = ""
  name_alias = ""
}
`
const testConfigMplsNodeSidPChildrenDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  loopback_address = "1.1.1.1"
  name = "node_sid_profile"
  segment_id = "1"
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

const testConfigMplsNodeSidPChildrenRemoveFromConfigDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  loopback_address = "1.1.1.1"
  name = "node_sid_profile"
  segment_id = "1"
}
`

const testConfigMplsNodeSidPChildrenRemoveOneDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  loopback_address = "1.1.1.1"
  name = "node_sid_profile"
  segment_id = "1"
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

const testConfigMplsNodeSidPChildrenRemoveAllDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  loopback_address = "1.1.1.1"
  name = "node_sid_profile"
  segment_id = "1"
  annotations = []
  tags = []
}
`
