
resource "aci_external_management_network_instance_profile" "full_example" {
  annotation  = "annotation"
  description = "description"
  name        = "test_name"
  name_alias  = "name_alias"
  priority    = "level1"
  relation_to_consumed_out_of_band_contracts = [
    {
      annotation                = "annotation_0"
      priority                  = "priority_0"
      out_of_band_contract_name = aci_out_of_band_contract.example.name
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


