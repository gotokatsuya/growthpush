# growthpush
  
GrowthPush API client library for Go.

## API Docs
https://growthbeat.github.io/api/growthpush/v3/

## Installation

```bash
$ go get github.com/gotokatsuya/growthpush
```

## Usage

```go
import (
    "github.com/gotokatsuya/growthpush/dispatcher"
    clientsSVC "github.com/gotokatsuya/growthpush/service/clients"
)

func CreateNewClient() error {
    var (
        applicationID = "xxx-xxx-xxx"
        clientID      = "xxx-xxx-xxx"
        credentialID  = "xxx-xxx-xxx"
    )
    client := dispatcher.NewClientWithParam(applicationID, clientID, credentialID)
    
    var (
        token       = "xxx-xxx-xxx" // device token
        os          = "xxx-xxx-xxx" // [android, ios]
        environment = "xxx-xxx-xxx" // [development, production]
    )
    _, err := clientsSVC.CreateNewClient(client, clientsSVC.CreateNewClientParameter{
        Token: token,
        OS: os,
        Environment: environment,
    })
    if err != nil {
        return err
    }
    return nil
}
```
