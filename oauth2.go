package akerun

import (
	"context"

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
