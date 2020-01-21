package apiclient

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/retry.v1"
)

func TestDelete(t *testing.T) {

	// handleDelete is a function used by a test server to respond to apiclient requests
	handleDelete := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() == "/v1/organisation/accounts/validAccountID%3Fversion=0" && req.Method == "DELETE" {
			responseJSON := `{}`
			rw.WriteHeader(204)
			rw.Write([]byte(responseJSON))
		}

		if req.URL.String() == "/v1/organisation/accounts/notFoundAccount%3Fversion=0" && req.Method == "DELETE" {
			responseJSON := `{"error_message":"record bd27e265-9605-4b4b-a0e5-3003ea9cc4dc does not exist"}`
			rw.WriteHeader(404)
			rw.Write([]byte(responseJSON))
		}

		if req.URL.String() == "/v1/organisation/accounts/internalServerError%3Fversion=0" && req.Method == "DELETE" {
			responseJSON := `{"error_message":"internal server error"}`
			rw.WriteHeader(500)
			rw.Write([]byte(responseJSON))
		}

		if req.URL.String() == "/v1/organisation/accounts/missingVersion%3Fversion=0" && req.Method == "DELETE" {
			responseJSON := `{"error_message":"version missing"}`
			rw.WriteHeader(400)
			rw.Write([]byte(responseJSON))
		}

	}

	// create a new client and configure it to use test server instead of the real API endpoint
	client := New("http://localhost:8080", 10*time.Second)
	testServer := httptest.NewServer(http.HandlerFunc(handleDelete))
	client.BaseURL = testServer.URL

	//Shorten retry duration to prevent test timeout
	exp := retry.Exponential{
		Initial: 10 * time.Millisecond,
		Factor:  1.5,
		Jitter:  true,
	}
	strategy := retry.LimitTime(10*time.Millisecond, exp)
	client.RetryStrategy = strategy

	tests := []struct {
		accountID string
		version   int
		err       error
	}{
		{"validAccountID", 0, nil},
		{"notFoundAccount", 0, errors.New("Status Code Not OK")},
		{"internalServerError", 0, errors.New("Retry timeout error")},
	}

	for _, test := range tests {
		err := Delete(client, test.accountID, test.version)
		assert.Equal(t, test.err, err)
	}
}
