package growthpush

import (
	"testing"

	"github.com/gotokatsuya/growthpush/dispatcher"
	clientsSVC "github.com/gotokatsuya/growthpush/service/clients"
	tagClientsSVC "github.com/gotokatsuya/growthpush/service/tagclients"
)

func TestCreateNewClientThenCreateNewTag(t *testing.T) {
	var (
		applicationID = "xxx-xxx-xxx"
		clientID      = "xxx-xxx-xxx"
		credentialID  = "xxx-xxx-xxx"
	)
	client := dispatcher.NewClientWithParam(applicationID, clientID, credentialID)
	var (
		token = "xxx-xxx-xxx" // device token
	)
	var (
		os          = "xxx-xxx-xxx" // [android, ios]
		environment = "xxx-xxx-xxx" // [development, production]
	)
	if _, err := clientsSVC.CreateNewClient(client, clientsSVC.CreateNewClientParameter{
		Token:       token,
		OS:          os,
		Environment: environment,
	}); err != nil {
		t.Fatal(err)
	}
	var (
		name  = "xxx-xxx-xxx"
		value = "xxx-xxx-xxx"
	)
	if _, err := tagClientsSVC.CreateNewTagByDeviceToken(client, tagClientsSVC.CreateNewTagByDeviceTokenParameter{
		Token: token,
		Name:  name,
		Value: value,
	}); err != nil {
		t.Fatal(err)
	}
}
