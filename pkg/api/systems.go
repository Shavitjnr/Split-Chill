package api

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type SystemsApi struct{}


var (
	Systems = &SystemsApi{}
)


func (a *SystemsApi) VersionHandler(c *core.WebContext) (any, *errs.Error) {
	result := make(map[string]string)

	result["version"] = settings.Version
	result["commitHash"] = settings.CommitHash

	if settings.BuildTime != "" {
		result["buildTime"] = settings.BuildTime
	}

	return result, nil
}
