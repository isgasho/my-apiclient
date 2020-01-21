package apiclient

import (
	"fmt"
)

// Delete deletes an account
func Delete(client *Client, accountID string, version int) error {

	path := fmt.Sprintf("/v1/organisation/accounts/%s?version=%d", accountID, version)

	_, err := client.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
