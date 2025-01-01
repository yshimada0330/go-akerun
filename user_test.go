package akerun

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestClient_GetUsers(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		assert.Equal(t, "/v3/organizations/org1/users", r.URL.Path)

		// Write a sample response
		_, err := w.Write([]byte(`{"users":[{"id":"user1"},{"id":"user2"}]}`))
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

	// Call the GetUsers method with some test parameters
	params := UsersParameter{Limit: 10}
	token := &oauth2.Token{AccessToken: "test_token"}
	users, err := client.GetUsers(context.Background(), token, "org1", params)

	// Check that the response was parsed correctly
	assert.NoError(t, err)
	assert.Len(t, users.Users, 2)
	assert.Equal(t, "user1", users.Users[0].ID)
	assert.Equal(t, "user2", users.Users[1].ID)
}

func TestClient_GetUser(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		assert.Equal(t, "/v3/organizations/org1/users/user1", r.URL.Path)

		// Write a sample response
		_, err := w.Write([]byte(`{"user":{"id":"user1","name":"Test User"}}`))
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

	// Call the GetUser method with some test parameters
	token := &oauth2.Token{AccessToken: "test_token"}
	user, err := client.GetUser(context.Background(), token, "org1", "user1")

	// Check that the response was parsed correctly
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.ID)
	assert.Equal(t, "Test User", user.Name)
}

func TestClient_RegisterUser(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		assert.Equal(t, "/v3/organizations/org1/users", r.URL.Path)

		// Write a sample response
		q := r.URL.Query()
		userName := q.Get("user_name")
		userEmail := q.Get("user_mail")
		_, err := w.Write([]byte(fmt.Sprintf(`{"user":{"id":"user1","name":"%s","mail":"%s"}}`, userName, userEmail)))
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

	// Call the RegisterUser method with some test parameters
	token := &oauth2.Token{AccessToken: "test_token"}
	param := RegisterUserParameter{
		UserMail: "example@example.com",
	}
	user, err := client.RegisterUser(context.Background(), token, "org1", "Test User", param)

	// Check that the response was parsed correctly
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.ID)
	assert.Equal(t, "Test User", user.Name)
}

func TestClient_UpdateUser(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		assert.Equal(t, "/v3/organizations/org1/users/user1", r.URL.Path)

		// Write a sample response
		q := r.URL.Query()
		userName := q.Get("user_name")
		userEmail := q.Get("user_mail")
		_, err := w.Write([]byte(fmt.Sprintf(`{"user":{"id":"user1","name":"%s","mail":"%s"}}`, userName, userEmail)))
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

	// Call the UpdateUser method with some test parameters
	token := &oauth2.Token{AccessToken: "test_token"}
	param := UpdateUserParameter{
		UserName: "Updated User",
		UserMail: "updated@example.com",
	}
	user, err := client.UpdateUser(context.Background(), token, "org1", "user1", param)

	// Check that the response was parsed correctly
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.ID)
	assert.Equal(t, "Updated User", user.Name)
}

func TestClient_ExitUser(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		assert.Equal(t, "/v3/organizations/org1/users/user1", r.URL.Path)
	}))
	defer ts.Close()

	os.Setenv("AKERUN_API_URL", ts.URL)

	// Create a new oauth2.Config with the test server's URL as the endpoint
	config := NewConfig("testId", "testPass", "http://localhost:8080/callback")

	// Create a new client with the oauth2.Config
	client := NewClient(config)

	// Call the ExitUser method with some test parameters
	token := &oauth2.Token{AccessToken: "test_token"}
	err := client.ExitUser(context.Background(), token, "org1", "user1")

	// Check that the response was parsed correctly
	assert.NoError(t, err)
}

func TestClient_InviteUser(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		assert.Equal(t, "/v3/organizations/org1/users/user1", r.URL.Path)

		// Write a sample response
		q := r.URL.Query()
		userId := q.Get("user_id")
		_, err := w.Write([]byte(fmt.Sprintf(`{"user":{"id":"%s","name":"Test User"}}`, userId)))
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

	// Call the InviteUser method with some test parameters
	token := &oauth2.Token{AccessToken: "test_token"}
	param := InviteUserParameter{
		UserAuthority: "admin",
	}
	user, err := client.InviteUser(context.Background(), token, "org1", "user1", param)

	// Check that the response was parsed correctly
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.ID)
	assert.Equal(t, "Test User", user.Name)
}
