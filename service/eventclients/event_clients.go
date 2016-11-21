package eventclients

import (
	"encoding/json"

	"github.com/gotokatsuya/growthpush/dispatcher"
	"github.com/gotokatsuya/growthpush/util"
)

const endpoint = "event_clients"

type CreateNewEventClientRequest struct {
	ClientID string `json:"clientId"`
	EventID  string `json:"eventId"`
	Value    string `json:"value"`
}

type CreateNewEventClientResponse struct {
	EventID  json.Number `json:"eventId"`
	ClientID json.Number `json:"clientId"`
}

func CreateNewEventClient(client *dispatcher.Client, req CreateNewEventClientRequest) (*CreateNewEventClientResponse, error) {
	parameters, err := util.JSONToMapString(req)
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchPostRequest(endpoint, parameters)
	if err != nil {
		return nil, err
	}
	resp := new(CreateNewEventClientResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
