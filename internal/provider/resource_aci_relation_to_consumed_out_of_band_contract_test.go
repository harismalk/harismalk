// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceMgmtRsOoBConsWithMgmtInstP(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigMgmtRsOoBConsMinDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "out_of_band_contract_name", "test_tn_vz_oob_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "priority", "unspecified"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigMgmtRsOoBConsAllDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "out_of_band_contract_name", "test_tn_vz_oob_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "priority", "level1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigMgmtRsOoBConsMinDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "out_of_band_contract_name", "test_tn_vz_oob_br_cp_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigMgmtRsOoBConsResetDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "out_of_band_contract_name", "test_tn_vz_oob_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "priority", "unspecified"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_consumed_out_of_band_contract.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "out_of_band_contract_name", "test_tn_vz_oob_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "priority", "unspecified"),
				),
			},
			// Update with children
			{
				Config:             testConfigMgmtRsOoBConsChildrenDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "out_of_band_contract_name", "test_tn_vz_oob_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "priority", "unspecified"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.0.key", "annotations_1"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.1.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_consumed_out_of_band_contract.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "out_of_band_contract_name", "test_tn_vz_oob_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "priority", "unspecified"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.0.key", "annotations_1"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.1.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigMgmtRsOoBConsChildrenRemoveFromConfigDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.0.key", "annotations_1"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.1.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigMgmtRsOoBConsChildrenRemoveOneDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.0.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigMgmtRsOoBConsChildrenRemoveAllDependencyWithMgmtInstP,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_consumed_out_of_band_contract.test", "annotations.#", "0"),
				),
			},
		},
	})
}

const testConfigMgmtRsOoBConsMinDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_relation_to_consumed_out_of_band_contract" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  out_of_band_contract_name = "test_tn_vz_oob_br_cp_name"
}
`

const testConfigMgmtRsOoBConsAllDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_relation_to_consumed_out_of_band_contract" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  out_of_band_contract_name = "test_tn_vz_oob_br_cp_name"
  annotation = "annotation"
  priority = "level1"
}
`

const testConfigMgmtRsOoBConsResetDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_relation_to_consumed_out_of_band_contract" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  out_of_band_contract_name = "test_tn_vz_oob_br_cp_name"
  annotation = "orchestrator:terraform"
  priority = "unspecified"
}
`
const testConfigMgmtRsOoBConsChildrenDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_relation_to_consumed_out_of_band_contract" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  out_of_band_contract_name = "test_tn_vz_oob_br_cp_name"
  annotations = [
	{
	  key = "annotations_1"
	  value = "value_1"
	},
	{
	  key = "annotations_2"
	  value = "value_2"
	},
  ]
}
`

const testConfigMgmtRsOoBConsChildrenRemoveFromConfigDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_relation_to_consumed_out_of_band_contract" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  out_of_band_contract_name = "test_tn_vz_oob_br_cp_name"
}
`

const testConfigMgmtRsOoBConsChildrenRemoveOneDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_relation_to_consumed_out_of_band_contract" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  out_of_band_contract_name = "test_tn_vz_oob_br_cp_name"
  annotations = [ 
	{
	  key = "annotations_2"
	  value = "value_2"
	},
  ]
}
`

const testConfigMgmtRsOoBConsChildrenRemoveAllDependencyWithMgmtInstP = testConfigMgmtInstPMin + `
resource "aci_relation_to_consumed_out_of_band_contract" "test" {
  parent_dn = aci_external_management_network_instance_profile.test.id
  out_of_band_contract_name = "test_tn_vz_oob_br_cp_name"
  annotations = []
}
`
