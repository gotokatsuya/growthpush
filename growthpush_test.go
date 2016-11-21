package growthpush

import (
	"flag"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/gotokatsuya/growthpush/dispatcher"
	"github.com/gotokatsuya/growthpush/service/clients"
	"github.com/gotokatsuya/growthpush/service/events"
	"github.com/gotokatsuya/growthpush/service/tagclients"
	"github.com/gotokatsuya/growthpush/service/tags"
)

var (
	deviceToken string
	deviceOS    string
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&deviceToken, "token", "", "")
	flag.StringVar(&deviceOS, "os", "", "")

	flag.Parse()
}

func getRandomName() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func mustGetClientID(t *testing.T, gpDispatcherClient *dispatcher.Client) string {
	getClientRes, getClientErr := clients.GetClientByToken(gpDispatcherClient, clients.GetClientByTokenRequest{
		Token: deviceToken,
	})
	if getClientRes != nil {
		t.Log("Exist device token")
		return getClientRes.ID.String()
	}

	t.Log(getClientErr.Error())

	newClientRes, newClientErr := clients.CreateNewClient(gpDispatcherClient, clients.CreateNewClientRequest{
		Token:       deviceToken,
		OS:          deviceOS,
		Environment: "development",
	})
	if newClientErr != nil {
		t.Fatal(newClientErr)
	}

	t.Log("Create device token")
	return newClientRes.ID.String()
}

func TestCreateTagClientFlow(t *testing.T) {
	var (
		applicationID = os.Getenv("GP_APPLICATION_ID")
		credentialID  = os.Getenv("GP_CREDENTIAL_ID")
	)
	gpDispatcherClient := dispatcher.NewClientWithParam(applicationID, credentialID)
	newTagRes, newTagErr := tags.CreateNewTag(gpDispatcherClient, tags.CreateNewTagRequest{
		Name: getRandomName(),
	})
	if newTagErr != nil {
		t.Fatal(newTagErr)
	}

	clientID := mustGetClientID(t, gpDispatcherClient)

	newTagClientRes, newTagClientErr := tagclients.CreateNewTagClient(gpDispatcherClient, tagclients.CreateNewTagClientRequest{
		ClientID: clientID,
		TagID:    newTagRes.ID.String(),
	})
	if newTagClientErr != nil {
		t.Fatal(newTagClientErr)
	}

	t.Log(newTagClientRes)
}

func TestCreateEventFlow(t *testing.T) {
	var (
		applicationID = os.Getenv("GP_APPLICATION_ID")
		credentialID  = os.Getenv("GP_CREDENTIAL_ID")
	)
	gpDispatcherClient := dispatcher.NewClientWithParam(applicationID, credentialID)
	newEventRes, newEventErr := events.CreateNewEvent(gpDispatcherClient, events.CreateNewEventRequest{
		Name: getRandomName(),
	})
	if newEventErr != nil {
		t.Fatal(newEventErr)
	}

	t.Log(newEventRes)
}
