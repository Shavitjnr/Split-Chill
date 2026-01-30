package api

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type DefaultApi struct{}


var (
	Default = &DefaultApi{}
)


func (a *DefaultApi) ApiNotFound(c *core.WebContext) (any, *errs.Error) {
	return nil, errs.ErrApiNotFound
}


func (a *DefaultApi) MethodNotAllowed(c *core.WebContext) (any, *errs.Error) {
	return nil, errs.ErrMethodNotAllowed
}
