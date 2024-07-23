// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceTagTagWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigTagTagDataSourceDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_tag.test", "key", "test_key"),
					resource.TestCheckResourceAttr("data.aci_tag.test", "value", "test_value"),
				),
			},
			{
				Config:      testConfigTagTagNotExistingFvTenant,
				ExpectError: regexp.MustCompile("Failed to read aci_tag data source"),
			},
		},
	})
}
func TestAccDataSourceTagTagWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.2(1l)") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigTagTagDataSourceDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_tag.test", "key", "test_key"),
					resource.TestCheckResourceAttr("data.aci_tag.test", "value", "test_value"),
				),
			},
			{
				Config:      testConfigTagTagNotExistingFvAEPg,
				ExpectError: regexp.MustCompile("Failed to read aci_tag data source"),
			},
		},
	})
}

const testConfigTagTagDataSourceDependencyWithFvTenant = testConfigTagTagMinDependencyWithFvTenant + `
data "aci_tag" "test" {
  parent_dn = aci_tenant.test.id
  key = "test_key"
  depends_on = [aci_tag.test]
}
`

const testConfigTagTagNotExistingFvTenant = testConfigTagTagMinDependencyWithFvTenant + `
data "aci_tag" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  key = "non_existing_key"
  depends_on = [aci_tag.test]
}
`
const testConfigTagTagDataSourceDependencyWithFvAEPg = testConfigTagTagMinDependencyWithFvAEPg + `
data "aci_tag" "test" {
  parent_dn = aci_application_epg.test.id
  key = "test_key"
  depends_on = [aci_tag.test]
}
`

const testConfigTagTagNotExistingFvAEPg = testConfigTagTagMinDependencyWithFvAEPg + `
data "aci_tag" "test_non_existing" {
  parent_dn = aci_application_epg.test.id
  key = "non_existing_key"
  depends_on = [aci_tag.test]
}
`
