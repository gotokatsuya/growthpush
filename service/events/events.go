package events

import (
	"encoding/json"

	"github.com/gotokatsuya/growthpush/dispatcher"
	"github.com/gotokatsuya/growthpush/util"
)

const endpoint = "events"

type CreateNewEventRequest struct {
	Name string `json:"name"`
}

type CreateNewEventResponse struct {
	ID json.Number `json:"id"`
}

func CreateNewEvent(client *dispatcher.Client, req CreateNewEventRequest) (*CreateNewEventResponse, error) {
	parameters, err := util.JSONToMapString(req)
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchPostRequest(endpoint, parameters)
	if err != nil {
		return nil, err
	}
	resp := new(CreateNewEventResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
