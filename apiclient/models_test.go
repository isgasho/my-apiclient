package apiclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountDataUnmarshalJSON(t *testing.T) {
	rawJSON := `{` +
		`"data":{` +
		`"type":"A",` +
		`"id":"B",` +
		`"organisation_id":"C",` +
		`"Attributes":{"country":"D",` +
		`"base_currency":"E",` +
		`"account_number":"F",` +
		`"bank_id":"G",` +
		`"bank_id_code":"H",` +
		`"bic":"I",` +
		`"iban":"J",` +
		`"title":"K",` +
		`"first_name":"L",` +
		`"bank_account_name":"M",` +
		`"alternative_bank_account_names":["N","O"],` +
		`"account_classification":"P",` +
		`"joint_account":true,` +
		`"account_matching_opt_out":true,` +
		`"secondary_identification":"Q"` +
		`}}}`

	expectedAccountData := AccountData{
		Data: Account{
			AccountType:    "A",
			ID:             "B",
			OrganisationID: "C",
			Attributes: AccountAttributes{
				Country:                     "D",
				BaseCurrency:                "E",
				AccountNumber:               "F",
				BankID:                      "G",
				BankIDCode:                  "H",
				Bic:                         "I",
				Iban:                        "J",
				Title:                       "K",
				FirstName:                   "L",
				BankAccountName:             "M",
				AlternativeBankAccountNames: []string{"N", "O"},
				AccountClassification:       "P",
				JointAccount:                true,
				AccountMatchingOptOut:       true,
				SecondaryIdentification:     "Q",
			},
		},
	}

	var accountData AccountData
	err := json.Unmarshal([]byte(rawJSON), &accountData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedAccountData, accountData)
}

func TestAccountDataMarshalJson(t *testing.T) {
	AccountData := AccountData{
		Data: Account{
			AccountType:    "A",
			ID:             "B",
			OrganisationID: "C",
			Attributes: AccountAttributes{
				Country:                     "D",
				BaseCurrency:                "E",
				AccountNumber:               "F",
				BankID:                      "G",
				BankIDCode:                  "H",
				Bic:                         "I",
				Iban:                        "J",
				Title:                       "K",
				FirstName:                   "L",
				BankAccountName:             "M",
				AlternativeBankAccountNames: []string{"N", "O"},
				AccountClassification:       "P",
				JointAccount:                true,
				AccountMatchingOptOut:       true,
				SecondaryIdentification:     "Q",
			},
		},
	}

	expectedJSON := `{` +
		`"data":{` +
		`"type":"A",` +
		`"id":"B",` +
		`"organisation_id":"C",` +
		`"attributes":{"country":"D",` +
		`"base_currency":"E",` +
		`"account_number":"F",` +
		`"bank_id":"G",` +
		`"bank_id_code":"H",` +
		`"bic":"I",` +
		`"iban":"J",` +
		`"title":"K",` +
		`"first_name":"L",` +
		`"bank_account_name":"M",` +
		`"alternative_bank_account_names":["N","O"],` +
		`"account_classification":"P",` +
		`"joint_account":true,` +
		`"account_matching_opt_out":true,` +
		`"secondary_identification":"Q"` +
		`}}}`

	json, err := json.Marshal(&AccountData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedJSON, string(json))

}

func TestAccountListDataUnmarshalJSON(t *testing.T) {
	rawJSON := `{` +
		`"data":[{` +
		`"attributes":{` +
		`"account_classification":"Personal",` +
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
		`"title":"Ms"` +
		`},` +
		`"created_on":"2020-01-15T21:41:09.508Z",` +
		`"id":"bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
		`"modified_on":"2020-01-15T21:41:09.508Z",` +
		`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
		`"type":"accounts",` +
		`"version":0` +
		`},{` +
		`"attributes":{` +
		`"account_classification":"Personal",` +
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
		`"title":"Ms"` +
		`},` +
		`"created_on":"2020-01-16T20:01:25.633Z",` +
		`"id":"cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
		`"modified_on":"2020-01-16T20:01:25.633Z",` +
		`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
		`"type":"accounts",` +
		`"version":0` +
		`}],` +
		`"links":{` +
		`"first":"/v1/organisation/accounts?page%5Bnumber%5D=first\u0026page%5Bsize%5D=%02",` +
		`"last":"/v1/organisation/accounts?page%5Bnumber%5D=last\u0026page%5Bsize%5D=%02",` +
		`"self":"/v1/organisation/accounts?page%5Bnumber%5D=%00\u0026page%5Bsize%5D=%02"` +
		`}}`

	expectedAccountListData := AccountListData(
		AccountListData{
			Data: []Account{
				Account{
					AccountType:    "accounts",
					ID:             "bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
					OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
					Attributes: AccountAttributes{Country: "GB",
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
				Account{AccountType: "accounts",
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

	var accountListData AccountListData
	err := json.Unmarshal([]byte(rawJSON), &accountListData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedAccountListData, accountListData)
}

func TestAccountListDataMarshalJSON(t *testing.T) {
	accountListData := AccountListData{
		Data: []Account{
			{
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
			{
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
					SecondaryIdentification:     "A1B2C3D4"},
			},
		},
		Links: PageLinks{
			First: "/v1/organisation/accounts?page%5Bnumber%5D=first\u0026page%5Bsize%5D=%02",
			Last:  "/v1/organisation/accounts?page%5Bnumber%5D=last\u0026page%5Bsize%5D=%02",
			Self:  "/v1/organisation/accounts?page%5Bnumber%5D=%00\u0026page%5Bsize%5D=%02",
		},
	}

	expectedJSON := `{` +
		`"data":[{` +
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
		`"secondary_identification":"A1B2C3D4"}},` +
		`{` +
		`"type":"accounts",` +
		`"id":"cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
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
		`"secondary_identification":"A1B2C3D4"}}],` +
		`"Links":{` +
		`"first":"/v1/organisation/accounts?page%5Bnumber%5D=first\u0026page%5Bsize%5D=%02",` +
		`"last":"/v1/organisation/accounts?page%5Bnumber%5D=last\u0026page%5Bsize%5D=%02",` +
		`"self":"/v1/organisation/accounts?page%5Bnumber%5D=%00\u0026page%5Bsize%5D=%02"` +
		`}}`

	json, err := json.Marshal(&accountListData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedJSON, string(json))
}
