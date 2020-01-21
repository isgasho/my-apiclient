package apiclient

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/retry.v1"
)

// Client is the type used to interface with the Accounts API
type Client struct {
	BaseURL       string
	HTTPClient    *http.Client
	RetryStrategy retry.Strategy
}

// New creates a new instance of a Client
func New(baseURL string, timeout time.Duration) *Client {
	exp := retry.Exponential{
		Initial: 10 * time.Millisecond,
		Factor:  1.5,
		Jitter:  true,
	}
	strategy := retry.LimitTime(60*time.Second, exp)

	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		RetryStrategy: strategy,
	}
}

// DoRequest makes a request to the Accounts API and handles the response 
func (c *Client) DoRequest(method string, path string, params *ListParams, payload io.Reader) (body []byte, err error) {
	reqURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	reqURL.Path = path

	if params != nil {
		query := url.Values{}
		if params.PageNum != nil {
			query.Add("page[number]", string(*params.PageNum))
		}
		if params.PageSize != nil {
			query.Add("page[size]", string(*params.PageSize))
		}
		reqURL.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(method, reqURL.String(), payload)
	if err != nil {
		return nil, err
	}

	for r := retry.Start(c.RetryStrategy, nil); r.Next(); {
		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 429, 500, 503, 504:
			log.Printf("Response Status %d Retrying request", resp.StatusCode)
			continue

		case 200, 201:
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			return body, nil

		case 204:
			return nil, nil

		default:
			err = errors.New("Status Code Not OK")
			return nil, err
		}
	}

	err = errors.New("Retry timeout error")
	return nil, err
}
