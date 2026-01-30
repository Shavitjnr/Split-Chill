package api

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type HealthsApi struct{}


var (
	Healths = &HealthsApi{}
)


func (a *HealthsApi) HealthStatusHandler(c *core.WebContext) (any, *errs.Error) {
	result := make(map[string]string)

	result["version"] = settings.Version
	result["commit"] = settings.CommitHash
	result["status"] = "ok"

	return result, nil
}
