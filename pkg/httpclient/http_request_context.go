package httpclient

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
)

const (
	logHandleKey = "log_handler"
)


type HttpResponseLogHandlerFunc func([]byte)


type httpRequestContext struct {
	core.Context
	logHandler HttpResponseLogHandlerFunc
}


func (c *httpRequestContext) Value(key any) any {
	if key == logHandleKey {
		return c.logHandler
	}

	return c.Context.Value(key)
}


func CustomHttpResponseLog(c core.Context, responseLogHandler HttpResponseLogHandlerFunc) core.Context {
	return &httpRequestContext{
		Context:    c,
		logHandler: responseLogHandler,
	}
}
