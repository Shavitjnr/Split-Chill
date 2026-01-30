package core

import (
	"net/http/httputil"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type CliHandlerFunc func(*CliContext) error


type MiddlewareHandlerFunc func(*WebContext)


type RedirectHandlerFunc func(*WebContext) (string, *errs.Error)


type ApiHandlerFunc func(*WebContext) (any, *errs.Error)


type JSONRPCApiHandlerFunc func(*WebContext, *JSONRPCRequest) (any, *errs.Error)


type EventStreamApiHandlerFunc func(*WebContext) *errs.Error


type DataHandlerFunc func(*WebContext) ([]byte, string, *errs.Error)


type ImageHandlerFunc func(*WebContext) ([]byte, string, *errs.Error)


type ProxyHandlerFunc func(*WebContext) (*httputil.ReverseProxy, *errs.Error)
