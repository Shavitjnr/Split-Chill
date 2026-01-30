package oauth2

import (
	"net/http"

	"golang.org/x/oauth2"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
)


type OAuth2Context struct {
	core.Context
	httpClient *http.Client
}


func (c *OAuth2Context) Value(key any) any {
	if key == oauth2.HTTPClient {
		return c.httpClient
	}

	return c.Context.Value(key)
}

func wrapOAuth2Context(ctx core.Context, httpClient *http.Client) core.Context {
	return &OAuth2Context{
		Context:    ctx,
		httpClient: httpClient,
	}
}
