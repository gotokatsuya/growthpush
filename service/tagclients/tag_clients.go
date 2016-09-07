package tagclients

import "github.com/gotokatsuya/growthpush/dispatcher"

const endpointTagClients = "tags"

type CreateNewTagByDeviceTokenParameter struct {
	Token string
	Name  string
	Value string
}

func CreateNewTagByDeviceToken(client *dispatcher.Client, param CreateNewTagByDeviceTokenParameter) ([]byte, error) {
	parameters := make(map[string]string)
	parameters["token"] = param.Token
	parameters["name"] = param.Name
	parameters["value"] = param.Value
	body, err := client.DispatchPostRequest(endpointTagClients, parameters)
	if err != nil {
		return nil, err
	}
	return body, nil
}
