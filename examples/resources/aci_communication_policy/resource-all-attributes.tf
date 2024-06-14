
resource "aci_communication_policy" "full_example" {
  annotation                         = "annotation"
  description                        = "description"
  name                               = "test_name"
  name_alias                         = "name_alias"
  owner_key                          = "owner_key"
  owner_tag                          = "owner_tag"
  strict_security_on_apic_oob_subnet = "no"
  http_ssl_configuration = [
    {
      access_control_allow_credential = "access_control_allow_credential_0"
      access_control_allow_origins    = "access_control_allow_origins_0"
      admin_st                        = "admin_st_0"
      annotation                      = "annotation_0"
      cli_only_mode                   = "cli_only_mode_0"
      client_cert_auth_state          = "client_cert_auth_state_0"
      description                     = "description_0"
      dh_param                        = "dh_param_0"
      global_throttle_rate            = "global_throttle_rate_0"
      global_throttle_st              = "global_throttle_st_0"
      global_throttle_unit            = "global_throttle_unit_0"
      max_request_status_count        = "max_request_status_count_0"
      name                            = "name_0"
      name_alias                      = "name_alias_0"
      node_exporter                   = "node_exporter_0"
      port                            = "port_0"
      referer                         = "referer_0"
      server_header                   = "server_header_0"
      ssl_protocols                   = "ssl_protocols_0"
      throttle_rate                   = "throttle_rate_0"
      throttle_st                     = "throttle_st_0"
      visore_access                   = "visore_access_0"
      tp = [
        {
          annotation = "annotation_0"
          target_dn  = "target_dn_0"
          annotations = [
            {
              key   = "key_0"
              value = "value_0"
            }
          ]
          tags = [
            {
              key   = "key_0"
              value = "value_0"
            }
          ]
        }
      ]
      key_ring = [
        {
          annotation           = "annotation_0"
          tn_pki_key_ring_name = aci_.example.name
          annotations = [
            {
              key   = "key_0"
              value = "value_0"
            }
          ]
          tags = [
            {
              key   = "key_0"
              value = "value_0"
            }
          ]
        }
      ]
      annotations = [
        {
          key   = "key_0"
          value = "value_0"
        }
      ]
      tags = [
        {
          key   = "key_0"
          value = "value_0"
        }
      ]
    }
  ]
  annotations = [
    {
      key   = "key_0"
      value = "value_0"
    }
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_0"
    }
  ]
}


