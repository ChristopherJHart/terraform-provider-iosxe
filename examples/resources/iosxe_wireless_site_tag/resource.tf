resource "iosxe_wireless_site_tag" "example" {
  site_tag_name   = "ST1"
  description     = "My Site Tag"
  ap_join_profile = "default-ap-profile"
  flex_profile    = "default-flex-profile"
  local_site      = false
}
