package akerun

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// AuthCodeURL returns a URL to OAuth 2.0 provider's consent page that asks for permissions for the required scopes explicitly.
func (c *Client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.config.Oauth2.AuthCodeURL(state, opts...)
}

// Exchange converts an authorization code into a token.
func (c *Client) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return c.config.Oauth2.Exchange(ctx, code, opts...)
}

// RefreshToken returns a new token that carries the same authorization as token, but with a renewed access token.
func (c *Client) RefreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	return c.config.Oauth2.TokenSource(ctx, token).Token()
}

// Revoke revokes the specified OAuth2 token.
func (c *Client) Revoke(ctx context.Context, token *oauth2.Token) error {
	postBody := map[string]string{
		"client_id":     c.config.Oauth2.ClientID,
		"client_secret": c.config.Oauth2.ClientSecret,
		"token":         token.AccessToken,
	}

	err := c.call(ctx, "oauth/revoke", http.MethodPost, token, nil, postBody, nil)
	if err != nil {
		return err
	}

	return nil
}
