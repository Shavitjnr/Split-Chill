package settings

import (
	"fmt"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
)


type ConfigContainer struct {
	current *Config
}


var (
	Version    string
	CommitHash string
	BuildTime  string
	Container  = &ConfigContainer{}
)


func SetCurrentConfig(config *Config) {
	Container.current = config
}


func (c *ConfigContainer) GetCurrentConfig() *Config {
	return c.current
}

func GetUserAgent() string {
	if Version == "" {
		return core.ApplicationName
	}

	return fmt.Sprintf("%s/%s", core.ApplicationName, Version)
}
