package growthpush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	// API URL
	baseURL = "https://api.growthpush.com"
	// API Version
	version = "1"
)

// Client ...
type Client struct {
	APIPath          string
	GrowthPushConfig Config
}

// NewClient ...
func NewClient(apiPath string) Client {
	return Client{
		APIPath:          apiPath,
		GrowthPushConfig: NewGrowthPushConfig(),
	}
}

// Post Request by Post Method.
func (c *Client) Post(values url.Values) ([]byte, error) {

	// url
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(fmt.Sprintf("Method = %s, Error = %s ", "GrowthPush.Client.Post.Parse", err.Error()))
		return nil, err
	}
	u.Path = path.Join(version, c.APIPath)
	urlString := u.String()

	// create request obj
	req, err := http.NewRequest("POST", urlString, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println(fmt.Sprintf("Method = %s, Error = %s ", "GrowthPush.Client.Post.NewRequest", err.Error()))
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// do request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Sprintf("Method = %s, Error = %s ", "GrowthPush.Client.Post.Do", err.Error()))
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("StatusCode = %d, Error = %s ", resp.StatusCode, "Invalid Request.")
	}

	byteAry, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteAry, nil
}
