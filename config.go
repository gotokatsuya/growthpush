package growthpush

import (
	"github.com/BurntSushi/toml"
)

// Config is growthpush's setting.
type Config struct {
	ApplicationID string `toml:"application_id"`
	SecretKey     string `toml:"secret_key"`
	Production    bool   `toml:"production"`
}

// NewGrowthPushConfig is a new instance config of growthpush.
func NewGrowthPushConfig() Config {
	var config Config
	configTmlPath := "config.tml"
	toml.DecodeFile(configTmlPath, &config)
	return config
}
