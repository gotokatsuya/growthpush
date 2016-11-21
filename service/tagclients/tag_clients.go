package tagclients

import (
	"encoding/json"

	"github.com/gotokatsuya/growthpush/dispatcher"
	"github.com/gotokatsuya/growthpush/util"
)

const endpoint = "tag_clients"

type CreateNewTagClientRequest struct {
	ClientID string `json:"clientId"`
	TagID    string `json:"tagId"`
	Value    string `json:"value"`
}

type CreateNewTagClientResponse struct {
	TagID    json.Number `json:"tagId"`
	ClientID json.Number `json:"clientId"`
}

func CreateNewTagClient(client *dispatcher.Client, req CreateNewTagClientRequest) (*CreateNewTagClientResponse, error) {
	parameters, err := util.JSONToMapString(req)
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchPostRequest(endpoint, parameters)
	if err != nil {
		return nil, err
	}
	resp := new(CreateNewTagClientResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
