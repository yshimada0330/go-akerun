package akerun

import (
	"reflect"
	"testing"

	"golang.org/x/oauth2"
)

func TestNewConfig(t *testing.T) {
	clientID := "000000"
	clientSecret := "xxxxxxxxxxxxxxxx"
	redirectURL := "http://localhost:8080/callback"

	expectConfig := &Config{
		APIUrl: APIUrl,
		Oauth2: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:   Oauth2AuthURL,
				TokenURL:  Oauth2TokenURL,
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
	}
	config := NewConfig(clientID, clientSecret, redirectURL)

	if reflect.DeepEqual(config, expectConfig) == false {
		t.Errorf("NewConfig() = %#v, want %#v", config, expectConfig)
	}
}

func TestNewClient(t *testing.T) {
	config := NewConfig("000000", "xxxxxxxxxxxxxxxx", "http://localhost:8080/callback")
	expectClient := &Client{config: config}
	client := NewClient(config)

	if reflect.DeepEqual(client, expectClient) == false {
		t.Errorf("NewConfig() = %#v, want %#v", client, expectClient)
	}
}
