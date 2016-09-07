package clients

import "github.com/gotokatsuya/growthpush/dispatcher"

const endpointClients = "clients"

type CreateNewClientParameter struct {
	Token       string
	OS          string
	Environment string
}

func CreateNewClient(client *dispatcher.Client, param CreateNewClientParameter) ([]byte, error) {
	parameters := make(map[string]string)
	parameters["token"] = param.Token
	parameters["os"] = param.OS
	parameters["environment"] = param.Environment
	body, err := client.DispatchPostRequest(endpointClients, parameters)
	if err != nil {
		return nil, err
	}
	return body, nil
}
