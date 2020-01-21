package apiclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createHandler(rw http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	bodyStr := string(body)

	switch {
	case req.URL.String() == "/v1/organisation/accounts" && req.Method == "POST" && strings.Contains(bodyStr, `"id":"bd27e265-9605-4b4b-a0e5-3003ea9cc4dc"`):
		responseJSON := `{` +
			`"data":{` +
			`"type":"accounts",` +
			`"id":"bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
			`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
			`"attributes":{"country":"GB",` +
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

		rw.WriteHeader(201)
		rw.Write([]byte(responseJSON))

	case req.URL.String() == "/v1/organisation/accounts" && req.Method == "POST" && strings.Contains(bodyStr, `"id":""`):
		responseJSON := `{"error_message":"validation failure list:\n` +
			`validation failure list:\nvalidation failure list:\n` +
			`account_classification in body should be one of [Personal Business]\n` +
			`country in body should match '^[A-Z]{2}$'\n` +
			`id in body is required\n` +
			`organisation_id in body is required\n` +
			`type in body is required"}`

		rw.WriteHeader(400)
		rw.Write([]byte(responseJSON))
	}
}

func TestCreate(t *testing.T) {
	validPayload := &AccountData{
		Data: Account{
			AccountType:    "accounts",
			ID:             "bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
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
				SecondaryIdentification:     "A1B2C3D4",
			},
		},
	}

	expectedAccount := AccountData(
		AccountData{
			Data: Account{
				AccountType:    "accounts",
				ID:             "bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
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
		payload     *AccountData
		accountData *AccountData
		err         error
	}{
		{validPayload, &expectedAccount, nil},
		{&AccountData{}, nil, errors.New("status code not ok")},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(createHandler))

	limitTimeout := 10 * time.Millisecond
	clientTimeout := 10 * time.Second
	client := New(testServer.URL, limitTimeout, clientTimeout)

	for _, test := range tests {
		accountData, err := Create(client, test.payload)
		assert.Equal(t, test.accountData, accountData)
		assert.Equal(t, test.err, err)
	}
}
