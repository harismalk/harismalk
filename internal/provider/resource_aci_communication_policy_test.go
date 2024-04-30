// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceCommPol(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config: testConfigCommPolMin,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "false"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config: testConfigCommPolAll,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", "owner_tag"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "no"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config: testConfigCommPolMin,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", "owner_tag"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "no"),
				),
			},
			// Update with empty strings config or default value
			{
				Config: testConfigCommPolReset,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "false"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_communication_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "false"),
				),
			},
			// Update with children
			{
				Config: testConfigCommPolChildren,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "false"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.access_control_allow_credential", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.access_control_allow_origins", "access_control_allow_origins_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.admin_st", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.cli_only_mode", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.client_cert_auth_state", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.dh_param", "1024"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_rate", "global_throttle_rate_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_st", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_unit", "global_throttle_unit_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.max_request_status_count", "max_request_status_count_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.node_exporter", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.port", "port_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.referer", "referer_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.server_header", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.ssl_protocols", "TLSv1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.throttle_rate", "throttle_rate_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.throttle_st", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.visore_access", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.access_control_allow_credential", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.access_control_allow_origins", "access_control_allow_origins_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.admin_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.cli_only_mode", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.client_cert_auth_state", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.dh_param", "2048"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.global_throttle_rate", "global_throttle_rate_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.global_throttle_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.global_throttle_unit", "global_throttle_unit_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.max_request_status_count", "max_request_status_count_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.node_exporter", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.port", "port_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.referer", "referer_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.server_header", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.ssl_protocols", "TLSv1.1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.throttle_rate", "throttle_rate_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.throttle_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.visore_access", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.#", "2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.#", "2"),
				),
			},
			// Update with children removed from config
			{
				Config: testConfigCommPolChildrenRemoveFromConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "false"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.access_control_allow_credential", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.access_control_allow_origins", "access_control_allow_origins_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.admin_st", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.cli_only_mode", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.client_cert_auth_state", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.dh_param", "1024"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_rate", "global_throttle_rate_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_st", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_unit", "global_throttle_unit_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.max_request_status_count", "max_request_status_count_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.node_exporter", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.port", "port_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.referer", "referer_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.server_header", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.ssl_protocols", "TLSv1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.throttle_rate", "throttle_rate_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.throttle_st", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.visore_access", "disabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.access_control_allow_credential", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.access_control_allow_origins", "access_control_allow_origins_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.admin_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.cli_only_mode", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.client_cert_auth_state", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.dh_param", "2048"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.global_throttle_rate", "global_throttle_rate_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.global_throttle_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.global_throttle_unit", "global_throttle_unit_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.max_request_status_count", "max_request_status_count_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.node_exporter", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.port", "port_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.referer", "referer_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.server_header", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.ssl_protocols", "TLSv1.1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.throttle_rate", "throttle_rate_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.throttle_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.1.visore_access", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.#", "2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config: testConfigCommPolChildrenRemoveOne,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "false"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.access_control_allow_credential", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.access_control_allow_origins", "access_control_allow_origins_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.admin_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.cli_only_mode", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.client_cert_auth_state", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.description", "description_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.dh_param", "2048"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_rate", "global_throttle_rate_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.global_throttle_unit", "global_throttle_unit_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.max_request_status_count", "max_request_status_count_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.name", "name_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.node_exporter", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.port", "port_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.referer", "referer_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.server_header", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.ssl_protocols", "TLSv1.1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.throttle_rate", "throttle_rate_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.throttle_st", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.0.visore_access", "enabled"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config: testConfigCommPolChildrenRemoveAll,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "strict_security_on_apic_oob_subnet", "false"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "http_ssl_configuration.#", "0"),
					resource.TestCheckResourceAttr("aci_communication_policy.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigCommPolMin = `
resource "aci_communication_policy" "test" {
  name = "test_name"
}
`

const testConfigCommPolAll = `
resource "aci_communication_policy" "test" {
  name = "test_name"
  annotation = "annotation"
  description = "description"
  name_alias = "name_alias"
  owner_key = "owner_key"
  owner_tag = "owner_tag"
  strict_security_on_apic_oob_subnet = "no"
}
`

const testConfigCommPolReset = `
resource "aci_communication_policy" "test" {
  name = "test_name"
  annotation = "orchestrator:terraform"
  description = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  strict_security_on_apic_oob_subnet = "false"
}
`
const testConfigCommPolChildren = `
resource "aci_communication_policy" "test" {
  name = "test_name"
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
  http_ssl_configuration = [
	{
	  access_control_allow_credential = "disabled"
	  access_control_allow_origins = "access_control_allow_origins_1"
	  admin_st = "disabled"
	  annotation = "annotation_1"
	  cli_only_mode = "disabled"
	  client_cert_auth_state = "disabled"
	  description = "description_1"
	  dh_param = "1024"
	  global_throttle_rate = "global_throttle_rate_1"
	  global_throttle_st = "disabled"
	  global_throttle_unit = "global_throttle_unit_1"
	  max_request_status_count = "max_request_status_count_1"
	  name = "name_1"
	  name_alias = "name_alias_1"
	  node_exporter = "disabled"
	  port = "port_1"
	  referer = "referer_1"
	  server_header = "disabled"
	  ssl_protocols = "TLSv1"
	  throttle_rate = "throttle_rate_1"
	  throttle_st = "disabled"
	  visore_access = "disabled"
	},
	{
	  access_control_allow_credential = "enabled"
	  access_control_allow_origins = "access_control_allow_origins_2"
	  admin_st = "enabled"
	  annotation = "annotation_2"
	  cli_only_mode = "enabled"
	  client_cert_auth_state = "enabled"
	  description = "description_2"
	  dh_param = "2048"
	  global_throttle_rate = "global_throttle_rate_2"
	  global_throttle_st = "enabled"
	  global_throttle_unit = "global_throttle_unit_2"
	  max_request_status_count = "max_request_status_count_2"
	  name = "name_2"
	  name_alias = "name_alias_2"
	  node_exporter = "enabled"
	  port = "port_2"
	  referer = "referer_2"
	  server_header = "enabled"
	  ssl_protocols = "TLSv1.1"
	  throttle_rate = "throttle_rate_2"
	  throttle_st = "enabled"
	  visore_access = "enabled"
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

const testConfigCommPolChildrenRemoveFromConfig = `
resource "aci_communication_policy" "test" {
  name = "test_name"
}
`

const testConfigCommPolChildrenRemoveOne = `
resource "aci_communication_policy" "test" {
  name = "test_name"
  annotations = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  http_ssl_configuration = [ 
	{
	  access_control_allow_credential = "enabled"
	  access_control_allow_origins = "access_control_allow_origins_2"
	  admin_st = "enabled"
	  annotation = "annotation_2"
	  cli_only_mode = "enabled"
	  client_cert_auth_state = "enabled"
	  description = "description_2"
	  dh_param = "2048"
	  global_throttle_rate = "global_throttle_rate_2"
	  global_throttle_st = "enabled"
	  global_throttle_unit = "global_throttle_unit_2"
	  max_request_status_count = "max_request_status_count_2"
	  name = "name_2"
	  name_alias = "name_alias_2"
	  node_exporter = "enabled"
	  port = "port_2"
	  referer = "referer_2"
	  server_header = "enabled"
	  ssl_protocols = "TLSv1.1"
	  throttle_rate = "throttle_rate_2"
	  throttle_st = "enabled"
	  visore_access = "enabled"
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

const testConfigCommPolChildrenRemoveAll = `
resource "aci_communication_policy" "test" {
  name = "test_name"
  annotations = []
  http_ssl_configuration = []
  tags = []
}
`
