// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceNetflowRsMonitorToExporterWithNetflowMonitorPol(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigNetflowRsMonitorToExporterMinDependencyWithNetflowMonitorPolAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.allow_test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.allow_test_2", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.allow_test_2", "annotation", "orchestrator:terraform"),
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
				Config:      testConfigNetflowRsMonitorToExporterMinDependencyWithNetflowMonitorPolAllowExisting,
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
				Config:             testConfigNetflowRsMonitorToExporterMinDependencyWithNetflowMonitorPolAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.allow_test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.allow_test_2", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.allow_test_2", "annotation", "orchestrator:terraform"),
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
				Config:             testConfigNetflowRsMonitorToExporterMinDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigNetflowRsMonitorToExporterAllDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigNetflowRsMonitorToExporterMinDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigNetflowRsMonitorToExporterResetDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_netflow_exporter.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with children
			{
				Config:             testConfigNetflowRsMonitorToExporterChildrenDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.1.value", "test_value"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_netflow_exporter.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.1.value", "test_value"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigNetflowRsMonitorToExporterChildrenRemoveFromConfigDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigNetflowRsMonitorToExporterChildrenRemoveOneDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigNetflowRsMonitorToExporterChildrenRemoveAllDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_netflow_exporter.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigNetflowRsMonitorToExporterMinDependencyWithNetflowMonitorPolAllowExisting = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
resource "aci_relation_to_netflow_exporter" "allow_test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
}
resource "aci_relation_to_netflow_exporter" "allow_test_2" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
  depends_on = [aci_relation_to_netflow_exporter.allow_test]
}
`

const testConfigNetflowRsMonitorToExporterMinDependencyWithNetflowMonitorPol = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
resource "aci_relation_to_netflow_exporter" "test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
}
`

const testConfigNetflowRsMonitorToExporterAllDependencyWithNetflowMonitorPol = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
resource "aci_relation_to_netflow_exporter" "test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
  annotation = "annotation"
}
`

const testConfigNetflowRsMonitorToExporterResetDependencyWithNetflowMonitorPol = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
resource "aci_relation_to_netflow_exporter" "test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
  annotation = "orchestrator:terraform"
}
`
const testConfigNetflowRsMonitorToExporterChildrenDependencyWithNetflowMonitorPol = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
resource "aci_relation_to_netflow_exporter" "test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
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

const testConfigNetflowRsMonitorToExporterChildrenRemoveFromConfigDependencyWithNetflowMonitorPol = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
resource "aci_relation_to_netflow_exporter" "test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
}
`

const testConfigNetflowRsMonitorToExporterChildrenRemoveOneDependencyWithNetflowMonitorPol = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
resource "aci_relation_to_netflow_exporter" "test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
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

const testConfigNetflowRsMonitorToExporterChildrenRemoveAllDependencyWithNetflowMonitorPol = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
resource "aci_relation_to_netflow_exporter" "test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
  annotations = []
  tags = []
}
`
