package github

import (
	"encoding/json"
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

const githubOAuth2AuthUrl = "https://github.com/login/oauth/authorize"     
const githubOAuth2TokenUrl = "https://github.com/login/oauth/access_token" 
const githubUserProfileApiUrl = "https://api.github.com/user"              
const githubUserEmailApiUrl = "https://api.github.com/user/emails"         

var githubOAuth2Scopes = []string{"user:email"}

type githubUserProfileResponse struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type githubUserEmailsResponse struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}


type GithubOAuth2Provider struct {
	provider.OAuth2Provider
	oauth2Config *oauth2.Config
}


func (p *GithubOAuth2Provider) GetOAuth2AuthUrl(c core.Context, state string, opts ...oauth2.AuthCodeOption) (string, error) {
	return p.oauth2Config.AuthCodeURL(state, opts...), nil
}


func (p *GithubOAuth2Provider) GetOAuth2Token(c core.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.oauth2Config.Exchange(c, code, opts...)
}


func (p *GithubOAuth2Provider) GetUserInfo(c core.Context, oauth2Token *oauth2.Token) (*data.OAuth2UserInfo, error) {
	
	req, err := p.buildAPIRequest(githubUserProfileApiUrl)

	if err != nil {
		log.Errorf(c, "[github_oauth2_provider.GetUserInfo] failed to get user info request, because %s", err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	oauth2Client := oauth2.NewClient(c, oauth2.StaticTokenSource(oauth2Token))

	req = req.WithContext(httpclient.CustomHttpResponseLog(c, func(data []byte) {
		log.Debugf(c, "[github_oauth2_provider.GetUserInfo] user profile response is %s", data)
	}))

	resp, err := oauth2Client.Do(req)

	if err != nil {
		log.Errorf(c, "[github_oauth2_provider.GetUserInfo] failed to get user info response, because %s", err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Errorf(c, "[github_oauth2_provider.GetUserInfo] failed to get user info response, because response code is %d", resp.StatusCode)
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	userProfileResp, err := p.parseUserProfile(c, body)

	if err != nil {
		return nil, err
	}

	
	req, err = p.buildAPIRequest(githubUserEmailApiUrl)

	if err != nil {
		log.Errorf(c, "[github_oauth2_provider.GetUserInfo] failed to get user emails request, because %s", err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	req = req.WithContext(httpclient.CustomHttpResponseLog(c, func(data []byte) {
		log.Debugf(c, "[github_oauth2_provider.GetUserInfo] user emails response is %s", data)
	}))

	resp, err = oauth2Client.Do(req)

	if err != nil {
		log.Errorf(c, "[github_oauth2_provider.GetUserInfo] failed to get user emails response, because %s", err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Errorf(c, "[github_oauth2_provider.GetUserInfo] failed to get user emails response, because response code is %d", resp.StatusCode)
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	email, err := p.parsePrimaryEmail(c, body)

	if err != nil {
		return nil, err
	}

	return &data.OAuth2UserInfo{
		UserName: userProfileResp.Login,
		Email:    email,
		NickName: userProfileResp.Name,
	}, nil
}

func (p *GithubOAuth2Provider) parseUserProfile(c core.Context, body []byte) (*githubUserProfileResponse, error) {
	userProfileResp := &githubUserProfileResponse{}
	err := json.Unmarshal(body, &userProfileResp)

	if err != nil {
		log.Warnf(c, "[github_oauth2_provider.parseUserProfile] failed to parse user profile response body, because %s", err.Error())
		return nil, errs.ErrCannotRetrieveUserInfo
	}

	if userProfileResp.Login == "" {
		log.Warnf(c, "[github_oauth2_provider.parseUserProfile] invalid user profile response body")
		return nil, errs.ErrCannotRetrieveUserInfo
	}

	return userProfileResp, nil
}

func (p *GithubOAuth2Provider) parsePrimaryEmail(c core.Context, body []byte) (string, error) {
	emailsResp := make([]githubUserEmailsResponse, 0)
	err := json.Unmarshal(body, &emailsResp)

	if err != nil {
		log.Warnf(c, "[github_oauth2_provider.parsePrimaryEmail] failed to parse user emails response body, because %s", err.Error())
		return "", errs.ErrCannotRetrieveUserInfo
	}

	for _, emailEntry := range emailsResp {
		if emailEntry.Primary && emailEntry.Verified {
			return emailEntry.Email, nil
		}
	}

	return "", nil
}

func (p *GithubOAuth2Provider) buildAPIRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	return req, nil
}


func NewGithubOAuth2Provider(config *settings.Config, redirectUrl string) (provider.OAuth2Provider, error) {
	oauth2Config := &oauth2.Config{
		ClientID:     config.OAuth2ClientID,
		ClientSecret: config.OAuth2ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  githubOAuth2AuthUrl,
			TokenURL: githubOAuth2TokenUrl,
		},
		RedirectURL: redirectUrl,
		Scopes:      githubOAuth2Scopes,
	}

	return &GithubOAuth2Provider{
		oauth2Config: oauth2Config,
	}, nil
}
