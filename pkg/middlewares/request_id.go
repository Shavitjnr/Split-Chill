package middlewares

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/requestid"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)

const requestIdHeader = "X-Request-ID"


func RequestId(config *settings.Config) core.MiddlewareHandlerFunc {
	return func(c *core.WebContext) {
		requestId := requestid.Container.GenerateRequestId(c.ClientIP(), c.ClientPort())
		c.SetContextId(requestId)

		if config.EnableRequestIdHeader {
			c.Header(requestIdHeader, requestId)
		}

		c.Next()
	}
}
