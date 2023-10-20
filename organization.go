package akerun

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathOrganizations = "organizations"
)

// Organization represents an organization in Akerun API.
type Organization struct {
	ID string `json:"id"`
}

// Organizations represents a list of organizations in Akerun API.
type Organizations struct {
	Organizations []Organization `json:"organizations"`
}

// OrganizationDetail represents the detailed information of an organization.
type OrganizationDetail struct {
	Organization OrganizationRow `json:"organization"`
}

// OrganizationRow represents a row in the organization table.
type OrganizationRow struct {
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
	err = c.call_version(ctx, APIPathOrganizations, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrganization retrieves the details of an organization with the specified ID.
func (c *Client) GetOrganization(ctx context.Context, oauth2Token *oauth2.Token, id string) (*OrganizationDetail, error) {
	var result OrganizationDetail
	err := c.call_version(ctx, path.Join(APIPathOrganizations, id), http.MethodGet, oauth2Token, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
