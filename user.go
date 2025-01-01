package akerun

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const apiPathUsers = "users"

type UsersList struct {
	Users []User `json:"users"`
}

type NFC struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	ImageUrl  string `json:"image_url"`
	Authority string `json:"authority"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Nfcs      []NFC  `json:"nfcs"`
}

type userRow struct {
	User User `json:"user"`
}

type UsersParameter struct {
	Limit           uint32 `url:"limit,omitempty"`
	IdAfter         string `url:"id_after,omitempty"`
	IdBefore        string `url:"id_before,omitempty"`
	UserCode        string `url:"user_code,omitempty"`
	UserMail        string `url:"user_mail,omitempty"`
	IncludeDateTime bool   `url:"include_date_time,omitempty"`
}

func (c *Client) GetUsers(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	params UsersParameter,
) (*UsersList, error) {
	var result UsersList
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathUsers), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetUser(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	userId string,
) (*User, error) {
	var result userRow
	err := c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathUsers, userId), http.MethodGet, oauth2Token, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.User, nil
}

type RegisterUserParameter struct {
	UserMail      string `url:"user_mail,omitempty"`
	UserImage     string `url:"user_image,omitempty"`
	UserAuthority string `url:"user_authority,omitempty"`
	UserCode      string `url:"user_code,omitempty"`
}

func (c *Client) RegisterUser(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	name string,
	params RegisterUserParameter,
) (*User, error) {
	var result userRow
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	v.Add("user_name", name)
	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathUsers), http.MethodPost, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.User, nil
}

type InviteUserParameter struct {
	UserImage     string `url:"user_image,omitempty"`
	UserAuthority string `url:"user_authority,omitempty"`
	UserCode      string `url:"user_code,omitempty"`
}

func (c *Client) InviteUser(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	userId string,
	params InviteUserParameter,
) (*User, error) {
	var result userRow
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	v.Add("user_id", userId)
	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathUsers, userId), http.MethodPost, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.User, nil
}

type UpdateUserParameter struct {
	UserName      string `url:"user_name,omitempty"`
	UserMail      string `url:"user_mail,omitempty"`
	UserImage     string `url:"user_image,omitempty"`
	UserAuthority string `url:"user_authority,omitempty"`
	UserCode      string `url:"user_code,omitempty"`
}

func (c *Client) UpdateUser(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	userId string,
	params UpdateUserParameter,
) (*User, error) {
	var result userRow
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	v.Add("user_id", userId)
	err = c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathUsers, userId), http.MethodPut, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.User, nil
}

func (c *Client) ExitUser(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	organizationId string,
	userId string,
) error {

	err := c.callVersion(ctx, path.Join(apiPathOrganizations, organizationId, apiPathUsers, userId), http.MethodDelete, oauth2Token, nil, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
