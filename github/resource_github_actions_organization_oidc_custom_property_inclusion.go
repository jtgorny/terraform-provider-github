package github

import (
	"context"
	"fmt"
	"net/url"

	gh "github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type actionsOIDCCustomPropertyInclusion struct {
	CustomPropertyName string `json:"custom_property_name,omitempty"`
	PropertyName       string `json:"property_name,omitempty"`
}

func (i actionsOIDCCustomPropertyInclusion) name() string {
	if i.CustomPropertyName != "" {
		return i.CustomPropertyName
	}

	return i.PropertyName
}

func resourceGithubActionsOrganizationOIDCCustomPropertyInclusion() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsOrganizationOIDCCustomPropertyInclusionCreate,
		Read:   resourceGithubActionsOrganizationOIDCCustomPropertyInclusionRead,
		Delete: resourceGithubActionsOrganizationOIDCCustomPropertyInclusionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"custom_property_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The organization repository custom property to include in GitHub Actions OIDC tokens.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 75)),
			},
		},
	}
}

func resourceGithubActionsOrganizationOIDCCustomPropertyInclusionCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	if err := checkOrganization(meta); err != nil {
		return err
	}

	propertyName := d.Get("custom_property_name").(string)
	inclusions, err := listActionsOrganizationOIDCCustomPropertyInclusions(ctx, client, orgName)
	if err != nil {
		return err
	}
	for _, inclusion := range inclusions {
		if inclusion.name() == propertyName {
			d.SetId(propertyName)
			return resourceGithubActionsOrganizationOIDCCustomPropertyInclusionRead(d, meta)
		}
	}

	endpoint := fmt.Sprintf("orgs/%s/actions/oidc/customization/properties/repo", orgName)
	req, err := client.NewRequest(ctx, "POST", endpoint, &actionsOIDCCustomPropertyInclusion{
		CustomPropertyName: propertyName,
	})
	if err != nil {
		return err
	}
	if _, err = client.Do(req, nil); err != nil {
		return err
	}

	d.SetId(propertyName)
	return resourceGithubActionsOrganizationOIDCCustomPropertyInclusionRead(d, meta)
}

func resourceGithubActionsOrganizationOIDCCustomPropertyInclusionRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	if err := checkOrganization(meta); err != nil {
		return err
	}

	inclusions, err := listActionsOrganizationOIDCCustomPropertyInclusions(ctx, client, orgName)
	if err != nil {
		return err
	}
	for _, inclusion := range inclusions {
		if inclusion.name() == d.Id() {
			return d.Set("custom_property_name", inclusion.name())
		}
	}

	d.SetId("")
	return nil
}

func resourceGithubActionsOrganizationOIDCCustomPropertyInclusionDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	if err := checkOrganization(meta); err != nil {
		return err
	}

	endpoint := fmt.Sprintf(
		"orgs/%s/actions/oidc/customization/properties/repo/%s",
		orgName,
		url.PathEscape(d.Id()),
	)
	req, err := client.NewRequest(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req, nil)
	if err != nil && (resp == nil || resp.StatusCode != 404) {
		return err
	}

	return nil
}

func listActionsOrganizationOIDCCustomPropertyInclusions(
	ctx context.Context,
	client *gh.Client,
	orgName string,
) ([]actionsOIDCCustomPropertyInclusion, error) {
	endpoint := fmt.Sprintf("orgs/%s/actions/oidc/customization/properties/repo", orgName)
	req, err := client.NewRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var inclusions []actionsOIDCCustomPropertyInclusion
	if _, err = client.Do(req, &inclusions); err != nil {
		return nil, err
	}

	return inclusions, nil
}
