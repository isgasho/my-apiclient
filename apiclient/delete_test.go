package apiclient

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func deleteHandler(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.String() == "/v1/organisation/accounts/validAccountID%3Fversion=0" && req.Method == "DELETE":
		responseJSON := `{}`
		rw.WriteHeader(204)
		rw.Write([]byte(responseJSON))

	case req.URL.String() == "/v1/organisation/accounts/notFoundAccount%3Fversion=0" && req.Method == "DELETE":
		responseJSON := `{"error_message":"record bd27e265-9605-4b4b-a0e5-3003ea9cc4dc does not exist"}`
		rw.WriteHeader(404)
		rw.Write([]byte(responseJSON))

	case req.URL.String() == "/v1/organisation/accounts/internalServerError%3Fversion=0" && req.Method == "DELETE":
		responseJSON := `{"error_message":"internal server error"}`
		rw.WriteHeader(500)
		rw.Write([]byte(responseJSON))

	case req.URL.String() == "/v1/organisation/accounts/missingVersion%3Fversion=0" && req.Method == "DELETE":
		responseJSON := `{"error_message":"version missing"}`
		rw.WriteHeader(400)
		rw.Write([]byte(responseJSON))
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		accountID string
		version   int
		err       error
	}{
		{"validAccountID", 0, nil},
		{"notFoundAccount", 0, errors.New("status code not ok")},
		{"internalServerError", 0, errors.New("retry timeout error")},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(deleteHandler))

	limitTimeout := 10 * time.Millisecond
	clientTimeout := 10 * time.Second
	client := New(testServer.URL, limitTimeout, clientTimeout)

	for _, test := range tests {
		err := Delete(client, test.accountID, test.version)
		assert.Equal(t, test.err, err)
	}
}
