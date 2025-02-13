// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvRsProvWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRsProvDataSourceDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_provided_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("data.aci_relation_to_provided_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_relation_to_provided_contract.test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("data.aci_relation_to_provided_contract.test", "priority", "unspecified"),
				),
			},
			{
				Config:      testConfigFvRsProvNotExistingFvAEPg,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_provided_contract data source"),
			},
		},
	})
}
func TestAccDataSourceFvRsProvWithFvESg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRsProvDataSourceDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_provided_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("data.aci_relation_to_provided_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_relation_to_provided_contract.test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("data.aci_relation_to_provided_contract.test", "priority", "unspecified"),
				),
			},
			{
				Config:      testConfigFvRsProvNotExistingFvESg,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_provided_contract data source"),
			},
		},
	})
}

const testConfigFvRsProvDataSourceDependencyWithFvAEPg = testConfigFvRsProvMinDependencyWithFvAEPg + `
data "aci_relation_to_provided_contract" "test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
  depends_on = [aci_relation_to_provided_contract.test]
}
`

const testConfigFvRsProvNotExistingFvAEPg = testConfigFvRsProvMinDependencyWithFvAEPg + `
data "aci_relation_to_provided_contract" "test_non_existing" {
  parent_dn = aci_application_epg.test.id
  contract_name = "non_existing_tn_vz_br_cp_name"
  depends_on = [aci_relation_to_provided_contract.test]
}
`
const testConfigFvRsProvDataSourceDependencyWithFvESg = testConfigFvRsProvMinDependencyWithFvESg + `
data "aci_relation_to_provided_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
  depends_on = [aci_relation_to_provided_contract.test]
}
`

const testConfigFvRsProvNotExistingFvESg = testConfigFvRsProvMinDependencyWithFvESg + `
data "aci_relation_to_provided_contract" "test_non_existing" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "non_existing_tn_vz_br_cp_name"
  depends_on = [aci_relation_to_provided_contract.test]
}
`
