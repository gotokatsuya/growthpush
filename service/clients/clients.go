package clients

import (
	"encoding/json"
	"path"

	"github.com/gotokatsuya/growthpush/dispatcher"
	"github.com/gotokatsuya/growthpush/util"
)

const endpoint = "clients"

type CreateNewClientRequest struct {
	Token       string `json:"token"`
	OS          string `json:"os"`
	Environment string `json:"environment"`
}

type CreateNewClientResponse struct {
	ID json.Number `json:"id"`
}

func CreateNewClient(client *dispatcher.Client, req CreateNewClientRequest) (*CreateNewClientResponse, error) {
	parameters, err := util.JSONToMapString(req)
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchPostRequest(endpoint, parameters)
	if err != nil {
		return nil, err
	}
	resp := new(CreateNewClientResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetClientByTokenRequest struct {
	Token string
}

type GetClientByTokenResponse struct {
	ID json.Number `json:"id"`
}

func GetClientByToken(client *dispatcher.Client, req GetClientByTokenRequest) (*GetClientByTokenResponse, error) {
	parameters, err := util.JSONToMapString(req)
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchGetRequest(path.Join(endpoint, "token", req.Token), parameters)
	if err != nil {
		return nil, err
	}
	resp := new(GetClientByTokenResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetClientsRequest struct {
	Limit            string `json:"limit"`
	ExclusiveStartID string `json:"exclusiveStartId"`
}

type GetClientsResponse struct {
	ID          json.Number `json:"id"`
	Token       string      `json:"token"`
	Status      string      `json:"status"`
	OS          string      `json:"os"`
	Environment string      `json:"environment"`
}

func (r GetClientsResponse) InvalidOrInactive() bool {
	return r.Status == "invalid" || r.Status == "inactive"
}

func (r GetClientsResponse) Development() bool {
	return r.Environment == "development"
}

func GetClients(client *dispatcher.Client, req GetClientsRequest) ([]GetClientsResponse, error) {
	parameters, err := util.JSONToMapString(req)
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchGetRequest(endpoint, parameters)
	if err != nil {
		return nil, err
	}
	var respList []GetClientsResponse
	if err := json.Unmarshal(body, &respList); err != nil {
		return nil, err
	}
	return respList, nil
}
