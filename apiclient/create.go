package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Create registers an existing bank account or creates a new one.
func Create(client *Client, account *AccountData) (*AccountData, error) {
	jsonPayload, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/v1/organisation/accounts")

	body, err := client.DoRequest("POST", path, nil, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	var newAccount AccountData
	err = json.Unmarshal(body, &newAccount)
	if err != nil {
		return nil, err
	}

	return &newAccount, nil
}
