---
page_title: "github_actions_organization_oidc_custom_property_inclusion (Resource) - GitHub"
description: |-
  Includes an organization repository custom property in GitHub Actions OIDC tokens
---

# github_actions_organization_oidc_custom_property_inclusion (Resource)

This resource includes an organization-level repository custom property in the OpenID Connect tokens issued to GitHub Actions workflows. The custom property must already exist in the organization.

Included properties appear as claims prefixed with `repo_property_`. For example, the custom property `owner` appears as `repo_property_owner`.

## Example Usage

```terraform
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
```

## Argument Reference

The following argument is supported:

- `custom_property_name` - (Required) The organization repository custom property to include in GitHub Actions OIDC tokens.

## Import

This resource can be imported using the custom property name.

```shell
terraform import github_actions_organization_oidc_custom_property_inclusion.owner owner
```
