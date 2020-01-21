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

var ( // Global variables referenced by tests requiring pointers to ints
	zero = 0
	one  = 1
	two  = 2
)

func TestList(t *testing.T) {

	// handleList is a function used by a test server to respond to apiclient requests
	handleList := func(rw http.ResponseWriter, req *http.Request) {

		if req.URL.String() == "/v1/organisation/accounts?page%5Bnumber%5D=%00&page%5Bsize%5D=%02" && req.Method == "GET" {
			responseJSON := `{"data":[{"attributes":{"account_classification":"Personal",` +
				`"account_matching_opt_out":false,` +
				`"account_number":"41426819",` +
				`"alternative_bank_account_names":["Sam Holder"],` +
				`"bank_account_name":"Samantha Holder",` +
				`"bank_id":"400300",` +
				`"bank_id_code":"GBDSC",` +
				`"base_currency":"GBP",` +
				`"bic":"NWBKGB22",` +
				`"country":"GB",` +
				`"first_name":"Samantha",` +
				`"iban":"GB11NWBK40030041426819",` +
				`"joint_account":false,` +
				`"secondary_identification":"A1B2C3D4",` +
				`"title":"Ms"},` +
				`"created_on":"2020-01-15T21:41:09.508Z",` +
				`"id":"bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
				`"modified_on":"2020-01-15T21:41:09.508Z",` +
				`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
				`"type":"accounts",` +
				`"version":0},` +
				`{"attributes":{"account_classification":"Personal",` +
				`"account_matching_opt_out":false,` +
				`"account_number":"41426819",` +
				`"alternative_bank_account_names":["Sam Holder"],` +
				`"bank_account_name":"Samantha Holder",` +
				`"bank_id":"400300",` +
				`"bank_id_code":"GBDSC",` +
				`"base_currency":"GBP",` +
				`"bic":"NWBKGB22",` +
				`"country":"GB",` +
				`"first_name":"Samantha",` +
				`"iban":"GB11NWBK40030041426819",` +
				`"joint_account":false,` +
				`"secondary_identification":"A1B2C3D4",` +
				`"title":"Ms"},` +
				`"created_on":"2020-01-16T20:01:25.633Z",` +
				`"id":"cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
				`"modified_on":"2020-01-16T20:01:25.633Z",` +
				`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
				`"type":"accounts",` +
				`"version":0}],` +
				`"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first\u0026page%5Bsize%5D=%02",` +
				`"last":"/v1/organisation/accounts?page%5Bnumber%5D=last\u0026page%5Bsize%5D=%02",` +
				`"self":"/v1/organisation/accounts?page%5Bnumber%5D=%00\u0026page%5Bsize%5D=%02"}}`

			rw.WriteHeader(200)
			rw.Write([]byte(responseJSON))
		}

		if req.URL.String() == "/v1/organisation/accounts?page%5Bnumber%5D=%00&page%5Bsize%5D=%01" && req.Method == "GET" {
			responseJSON := `{"error_message":"bad request"}`

			rw.WriteHeader(400)
			rw.Write([]byte(responseJSON))
		}

		if req.URL.String() == "/v1/organisation/accounts?page%5Bnumber%5D=%00&page%5Bsize%5D=%00" && req.Method == "GET" {
			responseJSON := `{"error_message":"internal server error"}`

			rw.WriteHeader(500)
			rw.Write([]byte(responseJSON))
		}

	}

	expectedAccountList := AccountListData(
		AccountListData{
			Data: []Account{
				Account{
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
						SecondaryIdentification:     "A1B2C3D4"}},
				Account{
					AccountType:    "accounts",
					ID:             "cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
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
						SecondaryIdentification:     "A1B2C3D4"}}},
			Links: PageLinks{
				First: "/v1/organisation/accounts?page%5Bnumber%5D=first\u0026page%5Bsize%5D=%02",
				Last:  "/v1/organisation/accounts?page%5Bnumber%5D=last\u0026page%5Bsize%5D=%02",
				Self:  "/v1/organisation/accounts?page%5Bnumber%5D=%00\u0026page%5Bsize%5D=%02",
			}})

	// create a new client and configure it to use test server instead of the real API endpoint
	client := New("http://localhost:8080", 10*time.Second)
	testServer := httptest.NewServer(http.HandlerFunc(handleList))
	client.BaseURL = testServer.URL

	//Shorten retry duration to prevent test timeout
	exp := retry.Exponential{
		Initial: 10 * time.Millisecond,
		Factor:  1.5,
		Jitter:  true,
	}
	strategy := retry.LimitTime(10*time.Millisecond, exp)
	client.RetryStrategy = strategy

	validParams := ListParams{PageNum: &zero, PageSize: &two}
	badRequest := ListParams{PageNum: &zero, PageSize: &one}
	internalServerError := ListParams{PageNum: &zero, PageSize: &zero}

	tests := []struct {
		params          ListParams
		accountListData *AccountListData
		err             error
	}{
		{validParams, &expectedAccountList, nil},
		{badRequest, &AccountListData{Data: []Account(nil), Links: PageLinks{First: "", Last: "", Self: ""}}, errors.New("Status Code Not OK")},
		{internalServerError, &AccountListData{Data: []Account(nil), Links: PageLinks{First: "", Last: "", Self: ""}}, errors.New("Retry timeout error")},
	}

	for _, test := range tests {
		accountListData, err := List(client, &test.params)
		assert.Equal(t, test.accountListData, accountListData)
		assert.Equal(t, test.err, err)
	}
}
