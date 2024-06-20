// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvFBRMemberWithFvFBRGroup(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvFBRMemberMinDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name_alias", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvFBRMemberAllDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name", "name"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name_alias", "name_alias"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvFBRMemberMinDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "fallback_member", "2.2.2.3"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvFBRMemberResetDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name_alias", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_vrf_fallback_route_group_member.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name_alias", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigFvFBRMemberChildrenDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_vrf_fallback_route_group_member.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvFBRMemberChildrenRemoveFromConfigDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvFBRMemberChildrenRemoveOneDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.deletable_child", "true"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.deletable_child", "true"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvFBRMemberChildrenRemoveAllDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group_member.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvFBRMemberMinDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route_group_member" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  fallback_member = "2.2.2.3"
}
`

const testConfigFvFBRMemberAllDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route_group_member" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  fallback_member = "2.2.2.3"
  annotation = "annotation"
  description = "description"
  name = "name"
  name_alias = "name_alias"
}
`

const testConfigFvFBRMemberResetDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route_group_member" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  fallback_member = "2.2.2.3"
  annotation = "orchestrator:terraform"
  description = ""
  name = ""
  name_alias = ""
}
`
const testConfigFvFBRMemberChildrenDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route_group_member" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  fallback_member = "2.2.2.3"
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

const testConfigFvFBRMemberChildrenRemoveFromConfigDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route_group_member" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  fallback_member = "2.2.2.3"
}
`

const testConfigFvFBRMemberChildrenRemoveOneDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route_group_member" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  fallback_member = "2.2.2.3"
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

const testConfigFvFBRMemberChildrenRemoveAllDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route_group_member" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  fallback_member = "2.2.2.3"
  annotations = [
  ]
  tags = [
  ]
}
`
