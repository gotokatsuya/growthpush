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
	defaultAPIVersion = "3"
)

type (
	RequiredParameter struct {
		ApplicationID string
		ClientID      string
		CredentialID  string
	}
	Client struct {
		HTTPClient *http.Client

		APIBaseURL           string
		APIVersion           string
		APIRequiredParameter RequiredParameter
	}
)

func NewClientWithParam(applicationID, clientID, credentialID string) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		APIBaseURL: defaultAPIBaseURL,
		APIVersion: defaultAPIVersion,
		APIRequiredParameter: RequiredParameter{
			ApplicationID: applicationID,
			ClientID:      clientID,
			CredentialID:  credentialID,
		},
	}
}

func (c *Client) DispatchPostRequest(endpoint string, parameters map[string]string) ([]byte, error) {
	u, err := url.Parse(c.APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(c.APIVersion, endpoint)
	urlString := u.String()

	values := url.Values{}
	// Required parameters
	values.Set("applicationId", c.APIRequiredParameter.ApplicationID)
	values.Set("clientId", c.APIRequiredParameter.ClientID)
	values.Set("credentialId", c.APIRequiredParameter.CredentialID)

	// Optional parameters
	for key, param := range parameters {
		values.Set(key, param)
	}

	resp, err := c.HTTPClient.Post(urlString, "application/json", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Return error when status code less than 200 or equal more than 300
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("StatusCode = %d, Message = %s ", resp.StatusCode, string(body))
	}
	return body, nil
}
