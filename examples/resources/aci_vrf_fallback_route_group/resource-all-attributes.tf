
resource "aci_vrf_fallback_route_group" "full_example_vrf" {
  parent_dn   = aci_vrf.example.id
  annotation  = "annotation"
  description = "description"
  name        = "fallback_route_group"
  name_alias  = "name_alias"
  vrf_fallback_route_group_members = [
    {
      annotation      = "annotation_0"
      description     = "description_0"
      name            = "name_0"
      name_alias      = "name_alias_0"
      fallback_member = "2.2.2.2"
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


