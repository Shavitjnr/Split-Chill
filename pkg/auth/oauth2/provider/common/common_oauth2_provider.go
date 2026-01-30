package common

import (
	"io"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/Shavitjnr/split-chill-ai/pkg/auth/oauth2/data"
	"github.com/Shavitjnr/split-chill-ai/pkg/auth/oauth2/provider"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/httpclient"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type CommonOAuth2Provider struct {
	provider.OAuth2Provider
	oauth2Config *oauth2.Config
	dataSource   CommonOAuth2DataSource
}


type CommonOAuth2DataSource interface {
	
	GetAuthUrl() string

	
	GetTokenUrl() string

	
	GetUserInfoRequest() (*http.Request, error)

	
	GetScopes() []string

	
	ParseUserInfo(c core.Context, body []byte) (*data.OAuth2UserInfo, error)
}


func (p *CommonOAuth2Provider) GetOAuth2AuthUrl(c core.Context, state string, opts ...oauth2.AuthCodeOption) (string, error) {
	return p.oauth2Config.AuthCodeURL(state, opts...), nil
}


func (p *CommonOAuth2Provider) GetOAuth2Token(c core.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.oauth2Config.Exchange(c, code, opts...)
}


func (p *CommonOAuth2Provider) GetUserInfo(c core.Context, oauth2Token *oauth2.Token) (*data.OAuth2UserInfo, error) {
	req, err := p.dataSource.GetUserInfoRequest()

	if err != nil {
		log.Errorf(c, "[common_oauth2_provider.GetUserInfo] failed to get user info request, because %s", err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	oauth2Client := oauth2.NewClient(c, oauth2.StaticTokenSource(oauth2Token))

	req = req.WithContext(httpclient.CustomHttpResponseLog(c, func(data []byte) {
		log.Debugf(c, "[common_oauth2_provider.GetUserInfo] response is %s", data)
	}))

	resp, err := oauth2Client.Do(req)

	if err != nil {
		log.Errorf(c, "[common_oauth2_provider.GetUserInfo] failed to get user info response, because %s", err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Errorf(c, "[common_oauth2_provider.GetUserInfo] failed to get user info response, because response code is %d", resp.StatusCode)
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	return p.dataSource.ParseUserInfo(c, body)
}


func (p *CommonOAuth2Provider) GetDataSource() CommonOAuth2DataSource {
	return p.dataSource
}


func NewCommonOAuth2Provider(config *settings.Config, redirectUrl string, dataSource CommonOAuth2DataSource) *CommonOAuth2Provider {
	oauth2Config := &oauth2.Config{
		ClientID:     config.OAuth2ClientID,
		ClientSecret: config.OAuth2ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  dataSource.GetAuthUrl(),
			TokenURL: dataSource.GetTokenUrl(),
		},
		RedirectURL: redirectUrl,
		Scopes:      dataSource.GetScopes(),
	}

	return &CommonOAuth2Provider{
		oauth2Config: oauth2Config,
		dataSource:   dataSource,
	}
}
