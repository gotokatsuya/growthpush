package tags

import (
	"encoding/json"

	"github.com/gotokatsuya/growthpush/dispatcher"
	"github.com/gotokatsuya/growthpush/util"
)

const endpoint = "tags"

type CreateNewTagRequest struct {
	Name string `json:"name"`
}

type CreateNewTagResponse struct {
	ID json.Number `json:"id"`
}

func CreateNewTag(client *dispatcher.Client, req CreateNewTagRequest) (*CreateNewTagResponse, error) {
	parameters, err := util.JSONToMapString(req)
	if err != nil {
		return nil, err
	}
	body, err := client.DispatchPostRequest(endpoint, parameters)
	if err != nil {
		return nil, err
	}
	resp := new(CreateNewTagResponse)
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
