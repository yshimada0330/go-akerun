package akerun

import (
	"context"

	"golang.org/x/oauth2"
)

func (c *Client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.config.Oauth2.AuthCodeURL(state, opts...)
}

func (c *Client) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return c.config.Oauth2.Exchange(ctx, code, opts...)
}

func (c *Client) RefreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	return c.config.Oauth2.TokenSource(ctx, token).Token()
}
