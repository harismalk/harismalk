// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceTagAnnotationWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigTagAnnotationMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "test_value"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigTagAnnotationAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "value"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigTagAnnotationMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "test_value"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigTagAnnotationResetDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "test_value"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_annotation.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "test_value"),
				),
			},
		},
	})
}
func TestAccResourceTagAnnotationWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigTagAnnotationMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "test_value"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigTagAnnotationAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "value"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigTagAnnotationMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "test_value"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigTagAnnotationResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "test_value"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_annotation.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("aci_annotation.test", "value", "test_value"),
				),
			},
		},
	})
}

const testConfigTagAnnotationMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_annotation" "test" {
  parent_dn = aci_tenant.test.id
  key = "test_key"
  value = "test_value"
}
`

const testConfigTagAnnotationAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_annotation" "test" {
  parent_dn = aci_tenant.test.id
  key = "test_key"
  value = "value"
}
`

const testConfigTagAnnotationResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_annotation" "test" {
  parent_dn = aci_tenant.test.id
  key = "test_key"
  value = "test_value"
}
`

const testConfigTagAnnotationMinDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_annotation" "test" {
  parent_dn = aci_application_epg.test.id
  key = "test_key"
  value = "test_value"
}
`

const testConfigTagAnnotationAllDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_annotation" "test" {
  parent_dn = aci_application_epg.test.id
  key = "test_key"
  value = "value"
}
`

const testConfigTagAnnotationResetDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_annotation" "test" {
  parent_dn = aci_application_epg.test.id
  key = "test_key"
  value = "test_value"
}
`
