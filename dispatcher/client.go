package dispatcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	defaultAPIBaseURL = "https://api.growthpush.com"
	defaultAPIVersion = "4"
)

type (
	RequiredParameter struct {
		ApplicationID string `json:"applicationId"`
		CredentialID  string `json:"credentialId"`
	}
	Client struct {
		HTTPClient *http.Client

		APIBaseURL           string
		APIVersion           string
		APIRequiredParameter RequiredParameter
	}
)

func NewClientWithParam(applicationID, credentialID string) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,

		APIBaseURL: defaultAPIBaseURL,
		APIVersion: defaultAPIVersion,
		APIRequiredParameter: RequiredParameter{
			ApplicationID: applicationID,
			CredentialID:  credentialID,
		},
	}
}

// DispatchGetRequest GrowthPushへのGETリクエスト
func (c *Client) DispatchGetRequest(endpoint string, parameters map[string]string) ([]byte, error) {
	u, err := url.Parse(c.APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(c.APIVersion, endpoint)

	values := url.Values{}
	values.Set("applicationId", c.APIRequiredParameter.ApplicationID)
	values.Set("credentialId", c.APIRequiredParameter.CredentialID)

	for key, param := range parameters {
		values.Set(key, param)
	}

	resp, err := http.Get(u.String() + "?" + values.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Return error when status code less than 200 or equal more than 300
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("Dispatcher.Client.DispatchGetRequest: code:%d body:%s", resp.StatusCode, string(respBody))
	}
	return respBody, nil
}

// DispatchPostRequest GrowthPushへのPOSTリクエスト
func (c *Client) DispatchPostRequest(endpoint string, parameters map[string]string) ([]byte, error) {
	u, err := url.Parse(c.APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(c.APIVersion, endpoint)

	values := url.Values{}
	values.Set("applicationId", c.APIRequiredParameter.ApplicationID)
	values.Set("credentialId", c.APIRequiredParameter.CredentialID)

	for key, param := range parameters {
		values.Set(key, param)
	}

	resp, err := http.PostForm(u.String(), values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("Dispatcher.Client.DispatchPostRequest: code:%d body:%s", resp.StatusCode, string(respBody))
	}
	return respBody, nil
}
