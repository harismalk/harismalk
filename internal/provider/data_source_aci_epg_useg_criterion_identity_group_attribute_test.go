// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvIdGroupAttrWithFvCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvIdGroupAttrDataSourceDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_epg_useg_criterion_identity_group_attribute.test", "selector", "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_criterion_identity_group_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_criterion_identity_group_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_criterion_identity_group_attribute.test", "name", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_criterion_identity_group_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_criterion_identity_group_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_criterion_identity_group_attribute.test", "owner_tag", ""),
				),
			},
			{
				Config:      testConfigFvIdGroupAttrNotExistingFvCrtrn,
				ExpectError: regexp.MustCompile("Failed to read aci_epg_useg_criterion_identity_group_attribute data source"),
			},
		},
	})
}

const testConfigFvIdGroupAttrDataSourceDependencyWithFvCrtrn = testConfigFvIdGroupAttrMinDependencyWithFvCrtrn + `
data "aci_epg_useg_criterion_identity_group_attribute" "test" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"
  depends_on = [aci_epg_useg_criterion_identity_group_attribute.test]
}
`

const testConfigFvIdGroupAttrNotExistingFvCrtrn = testConfigFvCrtrnMin + `
data "aci_epg_useg_criterion_identity_group_attribute" "test_non_existing" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng_non_existing"
}
`
