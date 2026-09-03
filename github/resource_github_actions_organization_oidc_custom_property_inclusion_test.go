package github

import "testing"

func TestGithubActionsOrganizationOIDCCustomPropertyInclusionSchema(t *testing.T) {
	t.Parallel()

	resource := resourceGithubActionsOrganizationOIDCCustomPropertyInclusion()
	property := resource.Schema["custom_property_name"]
	if !property.Required || !property.ForceNew {
		t.Fatal("custom_property_name must be required and force replacement")
	}
}

func TestActionsOIDCCustomPropertyInclusionName(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		inclusion actionsOIDCCustomPropertyInclusion
		want      string
	}{
		"create response": {
			inclusion: actionsOIDCCustomPropertyInclusion{CustomPropertyName: "owner"},
			want:      "owner",
		},
		"list response": {
			inclusion: actionsOIDCCustomPropertyInclusion{PropertyName: "owner"},
			want:      "owner",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := test.inclusion.name(); got != test.want {
				t.Fatalf("name() = %q, want %q", got, test.want)
			}
		})
	}
}
