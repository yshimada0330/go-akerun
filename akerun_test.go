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

func TestClient_GetAkeruns(t *testing.T) {
	organizationId := "test_org_id"
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request has the correct path
		url := fmt.Sprintf("/v3/organizations/%s/akeruns", organizationId)
		assert.Equal(t, url, r.URL.Path)

		// Write a sample response
		_, err := w.Write([]byte(`{"akeruns":[{"id":"A1030000","name":"オーナー0の部屋1","image_url":null,"autolock":false,"open_door_alert":false,"open_door_alert_second":null,"push_button":false,"normal_sound_volume":null,"alert_sound_volume":null,"battery_percentage":null,"seconds_till_autolock":null,"lock_type":null,"autolock_off_schedule":null,"akerun_remote":{"id":"TG11160000"},"nfc_reader_inside":{"id":"N1030000","battery_percentage":null},"nfc_reader_outside":{"id":"N0030000","battery_percentage":null},"door_sensor":{"id":"W1030000","battery_percentage":null}}]}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	originalValue := os.Getenv("AKERUN_API_URL")
	os.Setenv("AKERUN_API_URL", ts.URL)

	// Create a new oauth2.Config with the test server's URL as the endpoint
	config := NewConfig("testId", "testPass", "http://localhost:8080/callback")

	// Create a new client with the oauth2.Config
	client := NewClient(config)

	// Call the GetAkeruns method with some test parameters
	params := AkerunListParameter{Limit: 10}
	token := &oauth2.Token{AccessToken: "test_token"}
	orgs, err := client.GetAkeruns(context.Background(), token, organizationId, params)

	// Check that the response was parsed correctly
	assert.NoError(t, err)
	assert.Len(t, orgs.Akeruns, 1)
	assert.Equal(t, "A1030000", orgs.Akeruns[0].ID)

	os.Setenv("AKERUN_API_URL", originalValue)
}
