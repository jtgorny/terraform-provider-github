resource "github_organization_custom_properties" "owner" {
  property_name      = "owner"
  value_type         = "single_select"
  required           = false
  allowed_values     = ["java", "javascript"]
  values_editable_by = "org_actors"
}

resource "github_actions_organization_oidc_custom_property_inclusion" "owner" {
  custom_property_name = github_organization_custom_properties.owner.property_name
}
