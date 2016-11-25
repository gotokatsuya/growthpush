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

type GetEventsRequest struct {
	Limit            string `json:"limit"`
	ExclusiveStartID string `json:"exclusiveStartId"`
}

type GetEventsResponse struct {
	ID   json.Number `json:"id"`
	Name string      `json:"name"`
}

func GetEvents(client *dispatcher.Client, req GetEventsRequest) ([]GetEventsResponse, error) {
	parameters, err := util.JSONToMapString(req)
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchGetRequest(endpoint, parameters)
	if err != nil {
		return nil, err
	}
	var respList []GetEventsResponse
	if err := json.Unmarshal(body, &respList); err != nil {
		return nil, err
	}
	return respList, nil
}
