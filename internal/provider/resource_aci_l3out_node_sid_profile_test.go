// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
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
				Config:             testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfPAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "loopback_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "loopback_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "name_alias", ""),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfPAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfPAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "loopback_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "loopback_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.allow_test_2", "name_alias", ""),
				),
			},
		},
	})

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
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", "node_sid_profile"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", "name_alias_1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigMplsNodeSidPResetDependencyWithL3extLoopBackIfP,
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
			// Import testing
			{
				ResourceName:      "aci_l3out_node_sid_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigMplsNodeSidPChildrenDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.value", "test_value"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_l3out_node_sid_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "segment_id", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "loopback_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.value", "test_value"),
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
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigMplsNodeSidPChildrenRemoveOneDependencyWithL3extLoopBackIfP,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_node_sid_profile.test", "tags.0.value", "test_value"),
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

const testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfPAllowExisting = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "allow_test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
}
resource "aci_l3out_node_sid_profile" "allow_test_2" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
  depends_on = [aci_l3out_node_sid_profile.allow_test]
}
`

const testConfigMplsNodeSidPMinDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
}
`

const testConfigMplsNodeSidPAllDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
  annotation = "annotation"
  description = "description_1"
  loopback_address = "1.1.1.1"
  name = "node_sid_profile"
  name_alias = "name_alias_1"
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
  segment_id = "1"
  annotations = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigMplsNodeSidPChildrenRemoveFromConfigDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
}
`

const testConfigMplsNodeSidPChildrenRemoveOneDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigMplsNodeSidPChildrenRemoveAllDependencyWithL3extLoopBackIfP = testConfigL3extLoopBackIfPMinDependencyWithL3extRsNodeL3OutAtt + `
resource "aci_l3out_node_sid_profile" "test" {
  parent_dn = aci_l3out_loopback_interface_profile.test.id
  segment_id = "1"
  annotations = []
  tags = []
}
`
