# growthpush
  
GrowthPush API client library for Go.

## API Docs
https://growthbeat.github.io/api/growthpush/v4/

## Installation

```bash
$ go get github.com/gotokatsuya/growthpush
```

## Usage

```go
import (
    "github.com/gotokatsuya/growthpush/dispatcher"
    "github.com/gotokatsuya/growthpush/service/events"
)

func CreateEvent(eventName string) error {
	var (
		applicationID = os.Getenv("GP_APPLICATION_ID")
		credentialID  = os.Getenv("GP_CREDENTIAL_ID")
	)
	client := dispatcher.NewClientWithParam(applicationID, credentialID)
	res, err := events.CreateNewEvent(client, events.CreateNewEventRequest{
		Name: eventName,
	})
	if err != nil {
		return err
	}
    
	log.Println(res)

    return nil
}
```
