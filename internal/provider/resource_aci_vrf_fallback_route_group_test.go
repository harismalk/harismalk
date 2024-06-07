// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvFBRGroupWithFvCtx(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvFBRGroupMinDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvFBRGroupAllDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", "name_alias"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvFBRGroupMinDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvFBRGroupResetDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_vrf_fallback_route_group.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigFvFBRGroupChildrenDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.fallback_member", "2.2.2.2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotations", "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.tags", "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_vrf_fallback_route_group.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.fallback_member", "2.2.2.2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotations", "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.tags", "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvFBRGroupChildrenRemoveFromConfigDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.fallback_member", "2.2.2.2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotations", "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.tags", "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvFBRGroupChildrenRemoveOneDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotations", "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.description", "description_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name", "name_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.tags", "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvFBRGroupChildrenRemoveAllDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.#", "0"),
				),
			},
		},
	})
}

const testConfigFvFBRGroupMinDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
}
`

const testConfigFvFBRGroupAllDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
  annotation = "annotation"
  description = "description"
  name_alias = "name_alias"
}
`

const testConfigFvFBRGroupResetDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
  annotation = "orchestrator:terraform"
  description = ""
  name_alias = ""
}
`
const testConfigFvFBRGroupChildrenDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
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
  vrf_fallback_route_group_members = [
	{
	  annotation = "annotation_1"
	  description = "description_1"
	  fallback_member = "2.2.2.2"
	  name = "name_1"
	  name_alias = "name_alias_1"
	},
	{
	  annotation = "annotation_2"
	  annotations = "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"
	  description = "description_2"
	  fallback_member = "2.2.2.3"
	  name = "name_2"
	  name_alias = "name_alias_2"
	  tags = "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"
	},
  ]
}
`

const testConfigFvFBRGroupChildrenRemoveFromConfigDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
}
`

const testConfigFvFBRGroupChildrenRemoveOneDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
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
  vrf_fallback_route_group_members = [ 
	{
	  annotation = "annotation_2"
	  annotations = "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"
	  description = "description_2"
	  fallback_member = "2.2.2.3"
	  name = "name_2"
	  name_alias = "name_alias_2"
	  tags = "[map[key:key_0 value:value_1] map[key:key_1 value:value_2]]"
	},
  ]
}
`

const testConfigFvFBRGroupChildrenRemoveAllDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
  annotations = []
  tags = []
  vrf_fallback_route_group_members = []
}
`
