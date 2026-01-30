package gitea

import (
	"encoding/json"
	"net/http"

	"github.com/Shavitjnr/split-chill-ai/pkg/auth/oauth2/data"
	"github.com/Shavitjnr/split-chill-ai/pkg/auth/oauth2/provider"
	"github.com/Shavitjnr/split-chill-ai/pkg/auth/oauth2/provider/common"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)

type giteaUserInfoResponse struct {
	Login    string `json:"login"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}


type GiteaOAuth2DataSource struct {
	common.CommonOAuth2DataSource
	baseUrl string
}


func (s *GiteaOAuth2DataSource) GetAuthUrl() string {
	
	return s.baseUrl + "login/oauth/authorize"
}


func (s *GiteaOAuth2DataSource) GetTokenUrl() string {
	
	return s.baseUrl + "login/oauth/access_token"
}


func (s *GiteaOAuth2DataSource) GetUserInfoRequest() (*http.Request, error) {
	
	req, err := http.NewRequest("GET", s.baseUrl+"api/v1/user", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	return req, nil
}


func (s *GiteaOAuth2DataSource) GetScopes() []string {
	return []string{"read:user"}
}


func (s *GiteaOAuth2DataSource) ParseUserInfo(c core.Context, body []byte) (*data.OAuth2UserInfo, error) {
	userInfoResp := &giteaUserInfoResponse{}
	err := json.Unmarshal(body, &userInfoResp)

	if err != nil {
		log.Warnf(c, "[gitea_oauth2_datasource.ParseUserInfo] failed to parse user profile response body, because %s", err.Error())
		return nil, errs.ErrCannotRetrieveUserInfo
	}

	if userInfoResp.Login == "" {
		log.Warnf(c, "[gitea_oauth2_datasource.ParseUserInfo] invalid user profile response body")
		return nil, errs.ErrCannotRetrieveUserInfo
	}

	return &data.OAuth2UserInfo{
		UserName: userInfoResp.Login,
		Email:    userInfoResp.Email,
		NickName: userInfoResp.FullName,
	}, nil
}


func NewGiteaOAuth2Provider(config *settings.Config, redirectUrl string) (provider.OAuth2Provider, error) {
	if len(config.OAuth2GiteaBaseUrl) < 1 {
		return nil, errs.ErrInvalidOAuth2Config
	}

	baseUrl := config.OAuth2GiteaBaseUrl

	if baseUrl[len(baseUrl)-1] != '/' {
		baseUrl += "/"
	}

	return common.NewCommonOAuth2Provider(config, redirectUrl, &GiteaOAuth2DataSource{
		baseUrl: baseUrl,
	}), nil
}
