package growthpush

import (
	"testing"
)

func TestNewGrowthPushConfig(t *testing.T) {
	config := NewGrowthPushConfig()
	if len(config.ApplicationID) <= 0 {
		t.Fail()
	}
	if len(config.SecretKey) <= 0 {
		t.Fail()
	}

	t.Log(config.ApplicationID + " : " + config.SecretKey)
}
