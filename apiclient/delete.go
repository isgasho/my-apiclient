package apiclient

import (
	"fmt"
)

// Delete deletes an account.
func Delete(client *Client, accountID string, version int) error {
	path := fmt.Sprintf("/v1/organisation/accounts/%s?version=%d", accountID, version)

	if _, err := client.DoRequest("DELETE", path, nil, nil); err != nil {
		return err
	}

	return nil
}
