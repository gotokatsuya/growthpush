package dispatcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
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

const (
	GetMethod  = "GET"
	PostMethod = "POST"
)

func (c *Client) dispatchRequest(method, endpoint string, parameters map[string]string) ([]byte, error) {
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

	var (
		req    *http.Request
		reqErr error
	)
	switch method {
	case GetMethod:
		req, reqErr = http.NewRequest(method, u.String()+"?"+values.Encode(), nil)
	case PostMethod:
		req, reqErr = http.NewRequest(method, u.String(), strings.NewReader(values.Encode()))
	default:
		reqErr = fmt.Errorf("Not support %s method", method)
	}
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("Dispatcher.Client.DispatchPostRequest endpoint:%s code:%d body:%s", endpoint, resp.StatusCode, string(respBody))
	}
	return respBody, nil
}

// DispatchGetRequest GET request
func (c *Client) DispatchGetRequest(endpoint string, parameters map[string]string) ([]byte, error) {
	return c.dispatchRequest(GetMethod, endpoint, parameters)
}

// DispatchPostRequest POST request
func (c *Client) DispatchPostRequest(endpoint string, parameters map[string]string) ([]byte, error) {
	return c.dispatchRequest(PostMethod, endpoint, parameters)
}
