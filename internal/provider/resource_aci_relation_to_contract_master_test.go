// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRsSecInheritedWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test_2", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test_2", "annotation", "orchestrator:terraform"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvRsSecInheritedMinDependencyWithFvAEPgAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test_2", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test_2", "annotation", "orchestrator:terraform"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsSecInheritedResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_contract_master.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with children
			{
				Config:             testConfigFvRsSecInheritedChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_contract_master.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveFromConfigDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveOneDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: resource.ComposeAggregateTestCheckFunc(
			testCheckResourceDestroy,
		),
	})
}
func TestAccResourceFvRsSecInheritedWithFvESg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvESgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test_2", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test_2", "annotation", "orchestrator:terraform"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvRsSecInheritedMinDependencyWithFvESgAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvESgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test_2", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test_2", "annotation", "orchestrator:terraform"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedAllDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsSecInheritedResetDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_contract_master.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with children
			{
				Config:             testConfigFvRsSecInheritedChildrenDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_contract_master.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveFromConfigDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveOneDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveAllDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: resource.ComposeAggregateTestCheckFunc(
			testCheckResourceDestroy,
		),
	})
}

const testDependencyConfigFvRsSecInherited = `
resource "aci_endpoint_security_group" "test_0" {
  application_profile_dn = aci_application_profile.test.id
  name = "esg_0"
}
resource "aci_endpoint_security_group" "test_1" {
  application_profile_dn = aci_application_profile.test.id
  name = "esg_1"
}
resource "aci_application_epg" "test_2" {
  application_profile_dn = aci_application_profile.test.id
  name = "epg_2"
}
resource "aci_application_epg" "test_3" {
  application_profile_dn = aci_application_profile.test.id
  name = "epg_3"
}
`

const testConfigFvRsSecInheritedMinDependencyWithFvAEPgAllowExisting = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
}
resource "aci_relation_to_contract_master" "test_2" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
  depends_on = [aci_relation_to_contract_master.test]
}
`

const testConfigFvRsSecInheritedMinDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
}
`

const testConfigFvRsSecInheritedAllDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
  annotation = "annotation"
}
`

const testConfigFvRsSecInheritedResetDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
  annotation = "orchestrator:terraform"
}
`
const testConfigFvRsSecInheritedChildrenDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
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

const testConfigFvRsSecInheritedChildrenRemoveFromConfigDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
}
`

const testConfigFvRsSecInheritedChildrenRemoveOneDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
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

const testConfigFvRsSecInheritedChildrenRemoveAllDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_2.id
  annotations = []
  tags = []
}
`

const testConfigFvRsSecInheritedMinDependencyWithFvESgAllowExisting = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
}
resource "aci_relation_to_contract_master" "test_2" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
  depends_on = [aci_relation_to_contract_master.test]
}
`

const testConfigFvRsSecInheritedMinDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
}
`

const testConfigFvRsSecInheritedAllDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
  annotation = "annotation"
}
`

const testConfigFvRsSecInheritedResetDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
  annotation = "orchestrator:terraform"
}
`
const testConfigFvRsSecInheritedChildrenDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
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

const testConfigFvRsSecInheritedChildrenRemoveFromConfigDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
}
`

const testConfigFvRsSecInheritedChildrenRemoveOneDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
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

const testConfigFvRsSecInheritedChildrenRemoveAllDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_0.id
  annotations = []
  tags = []
}
`
