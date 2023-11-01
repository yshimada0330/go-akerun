package akerun

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	apiPathOrganizations = "organizations"
)

// id represents an ID of an organization.
type id struct {
	ID string `json:"id"`
}

// Organizations represents a list of organizations in Akerun API.
type Organizations struct {
	Organizations []id `json:"organizations"`
}

// row represents a row in the organization table.
type row struct {
	Organization Organization `json:"organization"`
}

// Organization represents the detailed information of an organization.
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OrganizationsParams represents the parameters for GetOrganizations method.
type OrganizationsParams struct {
	Limit    uint32 `url:"limit,omitempty"`
	IdAfter  string `url:"id_after,omitempty"`
	IdBefore string `url:"id_before,omitempty"`
}

// GetOrganizations returns a list of organizations.
func (c *Client) GetOrganizations(
	ctx context.Context, oauth2Token *oauth2.Token, params OrganizationsParams) (*Organizations, error) {

	var result Organizations
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	err = c.callVersion(ctx, apiPathOrganizations, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrganization retrieves the details of an organization with the specified ID.
func (c *Client) GetOrganization(ctx context.Context, oauth2Token *oauth2.Token, id string) (*Organization, error) {
	var result row
	err := c.callVersion(ctx, path.Join(apiPathOrganizations, id), http.MethodGet, oauth2Token, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.Organization, nil
}
