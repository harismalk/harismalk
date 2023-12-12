
data "aci_l3out_redistribute_policy" "example_l3_outside" {
  parent_dn                  = aci_l3_outside.example.id
  source                     = "direct"
  route_control_profile_name = "test_tn_rtctrl_profile_name"
}
