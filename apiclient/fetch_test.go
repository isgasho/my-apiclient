package apiclient

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func fetchHandler(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.String() == "/v1/organisation/accounts/validAccountID" && req.Method == "GET":
		responseJSON := `{` +
			`"data":{` +
			`"type":"accounts",` +
			`"id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
			`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
			`"Attributes":{"country":"GB",` +
			`"base_currency":"GBP",` +
			`"account_number":"41426819",` +
			`"bank_id":"400300",` +
			`"bank_id_code":"GBDSC",` +
			`"bic":"NWBKGB22",` +
			`"iban":"GB11NWBK40030041426819",` +
			`"title":"Ms",` +
			`"first_name":"Samantha",` +
			`"bank_account_name":"Samantha Holder",` +
			`"alternative_bank_account_names":["Sam Holder"],` +
			`"account_classification":"Personal",` +
			`"joint_account":false,` +
			`"account_matching_opt_out":false,` +
			`"secondary_identification":"A1B2C3D4"` +
			`}}}`

		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))

	case req.URL.String() == "/v1/organisation/accounts/notFoundAccount" && req.Method == "GET":
		responseJSON := `{"error_message":"record bd27e265-9605-4b4b-a0e5-3003ea9cc4dc does not exist"}`

		rw.WriteHeader(404)
		rw.Write([]byte(responseJSON))

	case req.URL.String() == "/v1/organisation/accounts/internalServerError" && req.Method == "GET":
		responseJSON := `{"error_message":"internal server error"}`

		rw.WriteHeader(500)
		rw.Write([]byte(responseJSON))
	}
}

func TestFetch(t *testing.T) {
	expectedAccount := AccountData(
		AccountData{
			Data: Account{
				AccountType:    "accounts",
				ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
				Attributes: AccountAttributes{
					Country:                     "GB",
					BaseCurrency:                "GBP",
					AccountNumber:               "41426819",
					BankID:                      "400300",
					BankIDCode:                  "GBDSC",
					Bic:                         "NWBKGB22",
					Iban:                        "GB11NWBK40030041426819",
					Title:                       "Ms",
					FirstName:                   "Samantha",
					BankAccountName:             "Samantha Holder",
					AlternativeBankAccountNames: []string{"Sam Holder"},
					AccountClassification:       "Personal",
					JointAccount:                false,
					AccountMatchingOptOut:       false,
					SecondaryIdentification:     "A1B2C3D4"},
			},
		},
	)

	tests := []struct {
		accountID   string
		accountData *AccountData
		err         error
	}{
		{"validAccountID", &expectedAccount, nil},
		{"notFoundAccount", nil, errors.New("status code not ok")},
		{"internalServerError", nil, errors.New("retry timeout error")},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fetchHandler))

	limitTimeout := 10 * time.Millisecond
	clientTimeout := 10 * time.Second
	client := New(testServer.URL, limitTimeout, clientTimeout)

	for _, test := range tests {
		accountData, err := Fetch(client, test.accountID)
		assert.Equal(t, test.accountData, accountData)
		assert.Equal(t, test.err, err)
	}
}
