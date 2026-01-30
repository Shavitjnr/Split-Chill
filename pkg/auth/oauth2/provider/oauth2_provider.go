package provider

import (
	"golang.org/x/oauth2"

	"github.com/Shavitjnr/split-chill-ai/pkg/auth/oauth2/data"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
)


type OAuth2Provider interface {
	
	GetOAuth2AuthUrl(c core.Context, state string, opts ...oauth2.AuthCodeOption) (string, error)

	
	GetOAuth2Token(c core.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)

	
	GetUserInfo(c core.Context, oauth2Token *oauth2.Token) (*data.OAuth2UserInfo, error)
}
