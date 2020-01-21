package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Create registers an existing bank account or creates a new one
func Create(client *Client, payload *AccountData) (*AccountData, error) {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return &AccountData{}, err
	}

	path := fmt.Sprintf("/v1/organisation/accounts")

	body, err := client.DoRequest("POST", path, nil, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	var createdData AccountData

	err = json.Unmarshal(body, &createdData)
	if err != nil {
		return &AccountData{}, err
	}

	return &createdData, nil
}
