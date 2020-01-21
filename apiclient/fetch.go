package apiclient

import (
	"encoding/json"
	"fmt"
)

// Fetch gets a single account using the accountID.
func Fetch(client *Client, accountID string) (*AccountData, error) {
	path := fmt.Sprintf("/v1/organisation/accounts/%s", accountID)

	body, err := client.DoRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var account AccountData
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}

	return &account, nil
}
