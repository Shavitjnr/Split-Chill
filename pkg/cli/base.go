package cli

import "github.com/Shavitjnr/split-chill-ai/pkg/settings"


type CliUsingConfig struct {
	container *settings.ConfigContainer
}


func (l *CliUsingConfig) CurrentConfig() *settings.Config {
	return l.container.GetCurrentConfig()
}
