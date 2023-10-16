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

const (
	APIEndpoint    = "https://api.akerun.com"
	APIPath1       = "/v3"
	Oauth2AuthURL  = "https://api.akerun.com/oauth/authorize"
	Oauth2TokenURL = "https://api.akerun.com/oauth/token"
)

type Config struct {
	APIEndpoint string
	Oauth2      *oauth2.Config
}

type Error struct {
	StatusCode int
	RawError   string
}

func (e *Error) Error() string {
	return e.RawError
}

func NewConfig(clientID, clientSecret, redirectURL string) *Config {
	apiEndpoint := os.Getenv("AKERUN_API_ENDPOINT")
	if apiEndpoint != "" {
		apiEndpoint = APIEndpoint
	}

	oauth2AuthURL := os.Getenv("AKERUN_OAUTH2_AUTH_URL")
	if oauth2AuthURL != "" {
		oauth2AuthURL = Oauth2AuthURL
	}

	oauth2TokenURL := os.Getenv("AKERUN_OAUTH2_TOKEN_URL")
	if oauth2TokenURL != "" {
		oauth2TokenURL = Oauth2TokenURL
	}

	return &Config{
		APIEndpoint: apiEndpoint,
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

type Client struct {
	httpClient *http.Client
	config     *Config
}

func NewClient(config *Config) *Client {
	return &Client{config: config}
}

func (c *Client) call(
	ctx context.Context,
	apiPath string,
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
	req, err := c.newRequest(ctx, apiPath, method, contentType, queryParams, body)
	if err != nil {
		return err
	}

	return c.do(ctx, oauth2Token, req, res)
}

func (c *Client) newRequest(
	ctx context.Context,
	apiPath string,
	method string,
	contentType string,
	queryParams url.Values,
	body io.Reader,
) (*http.Request, error) {
	u, err := url.Parse(c.config.APIEndpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, APIPath1, apiPath)
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
		res := &Error{
			StatusCode: code,
			RawError:   string(byt),
		}

		return res
	}

	if res == nil {
		return nil
	}
	return json.NewDecoder(r).Decode(&res)
}
