package client

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestServer(t *testing.T, path string, response string) (*httptest.Server, *CieloClient) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), path)
		// Send response to be tested
		_, err := rw.Write([]byte(response))
		if err != nil {
			t.Fatal(err)
		}
	}))

	c := NewClient(uuid.Must(uuid.NewUUID()), "merchantKey", &TestLogger{})
	c.HTTPClient = server.Client()
	c.environment.ApiUrl = server.URL

	return server, c
}
