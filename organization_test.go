package akerun

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestClient_GetOrganizations(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		assert.Equal(t, "/v3/organizations", r.URL.Path)

		// Write a sample response
		_, err := w.Write([]byte(`{"organizations":[{"id":"org1"},{"id":"org2"}]}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	os.Setenv("AKERUN_API_URL", ts.URL)

	// Create a new oauth2.Config with the test server's URL as the endpoint
	config := NewConfig("testId", "testPass", "http://localhost:8080/callback")

	// Create a new client with the oauth2.Config
	client := NewClient(config)

	// Call the GetOrganizations method with some test parameters
	params := OrganizationsParams{Limit: 10}
	token := &oauth2.Token{AccessToken: "test_token"}
	orgs, err := client.GetOrganizations(context.Background(), token, params)

	// Check that the response was parsed correctly
	assert.NoError(t, err)
	assert.Len(t, orgs.Organizations, 2)
	assert.Equal(t, "org1", orgs.Organizations[0].ID)
	assert.Equal(t, "org2", orgs.Organizations[1].ID)
}

func TestClient_GetOrganization(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		assert.Equal(t, "/v3/organizations/org1", r.URL.Path)

		// Write a sample response
		_, err := w.Write([]byte(`{"organization":{"id":"org1","name":"Test Org"}}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	os.Setenv("AKERUN_API_URL", ts.URL)

	// Create a new oauth2.Config with the test server's URL as the endpoint
	config := NewConfig("testId", "testPass", "http://localhost:8080/callback")

	// Create a new client with the oauth2.Config
	client := NewClient(config)

	// Call the GetOrganization method with some test parameters
	token := &oauth2.Token{AccessToken: "test_token"}
	org, err := client.GetOrganization(context.Background(), token, "org1")

	// Check that the response was parsed correctly
	assert.NoError(t, err)
	assert.Equal(t, "org1", org.Organization.ID)
	assert.Equal(t, "Test Org", org.Organization.Name)
}
