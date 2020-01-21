package apiclient

import (
	"encoding/json"
)

// ListParams are optional parameters used to call the List endpoint
type ListParams struct {
	PageNum, PageSize *int
}

// List accepts optional parameters and lists all accounts
func List(client *Client, params *ListParams) (*AccountListData, error) {

	path := "/v1/organisation/accounts"

	body, err := client.DoRequest("GET", path, params, nil)
	if err != nil {
		return &AccountListData{}, err
	}

	var accountList AccountListData
	err = json.Unmarshal(body, &accountList)
	if err != nil {
		return &AccountListData{}, nil
	}

	return &accountList, nil
}
