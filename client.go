package akerun

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"golang.org/x/oauth2"
)

// APIUrl is the base URL for the Akerun API.
const APIUrl = "https://api.akerun.com"

// APIVerison is the version of the Akerun API.
const APIVerison = "/v3"

// Oauth2AuthURL is the URL for the Akerun OAuth2 authorization endpoint.
const Oauth2AuthURL = "https://api.akerun.com/oauth/authorize"

// Oauth2TokenURL is the URL for the Akerun OAuth2 token endpoint.
const Oauth2TokenURL = "https://api.akerun.com/oauth/token"

// Config represents the configuration for the Akerun client.
type Config struct {
	APIUrl string
	Oauth2 *oauth2.Config
}

// Error represents an error returned by the Akerun API.
type Error struct {
	StatusCode int
	RawError   string
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.RawError
}

// NewConfig creates a new configuration for the Akerun client.
func NewConfig(clientID, clientSecret, redirectURL string) *Config {
	apiUrl := os.Getenv("AKERUN_API_URL")
	if apiUrl == "" {
		apiUrl = APIUrl
	}

	oauth2AuthURL := os.Getenv("AKERUN_OAUTH2_AUTH_URL")
	if oauth2AuthURL == "" {
		oauth2AuthURL = Oauth2AuthURL
	}

	oauth2TokenURL := os.Getenv("AKERUN_OAUTH2_TOKEN_URL")
	if oauth2TokenURL == "" {
		oauth2TokenURL = Oauth2TokenURL
	}

	return &Config{
		APIUrl: apiUrl,
		Oauth2: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:   oauth2AuthURL,
				TokenURL:  oauth2TokenURL,
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
	}
}

// Client represents the Akerun client.
type Client struct {
	httpClient *http.Client
	config     *Config
}

// NewClient creates a new Akerun client.
func NewClient(config *Config) *Client {
	return &Client{config: config}
}

// call sends a request to the Akerun API.
func (c *Client) call(
	ctx context.Context,
	apiEndpoint string,
	method string,
	oauth2Token *oauth2.Token,
	queryParams url.Values,
	postBody interface{},
	res interface{},
) error {
	var (
		contentType string
		body        io.Reader
	)
	contentType = "application/json"
	jsonParams, err := json.Marshal(postBody)
	if err != nil {
		return err
	}

	body = bytes.NewBuffer(jsonParams)
	req, err := c.newRequest(ctx, apiEndpoint, method, contentType, queryParams, body)
	if err != nil {
		return err
	}

	return c.do(ctx, oauth2Token, req, res)
}

// newRequest creates a new HTTP request for the Akerun API.
func (c *Client) newRequest(
	ctx context.Context,
	apiEndpoint string,
	method string,
	contentType string,
	queryParams url.Values,
	body io.Reader,
) (*http.Request, error) {
	u, err := url.Parse(c.config.APIUrl)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, APIVerison, apiEndpoint)
	u.RawQuery = queryParams.Encode()
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	return req, nil
}

// do sends an HTTP request to the Akerun API.
func (c *Client) do(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	req *http.Request,
	res interface{},
) error {
	tokenSource := c.config.Oauth2.TokenSource(ctx, oauth2Token)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	response, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	var r io.Reader = response.Body

	code := response.StatusCode
	if code >= http.StatusBadRequest {
		byt, _ := io.ReadAll(r)
		return &Error{
			StatusCode: code,
			RawError:   string(byt),
		}
	}

	if res == nil {
		return nil
	}
	return json.NewDecoder(r).Decode(&res)
}
