package core

const USER_EXTERNAL_AUTH_TYPE_CATEOGRY_OAUTH2 = "oauth2"


type UserExternalAuthType string


const (
	USER_EXTERNAL_AUTH_TYPE_OAUTH2_OIDC      UserExternalAuthType = "oidc"
	USER_EXTERNAL_AUTH_TYPE_OAUTH2_NEXTCLOUD UserExternalAuthType = "nextcloud"
	USER_EXTERNAL_AUTH_TYPE_OAUTH2_GITEA     UserExternalAuthType = "gitea"
	USER_EXTERNAL_AUTH_TYPE_OAUTH2_GITHUB    UserExternalAuthType = "github"
)


func (t UserExternalAuthType) GetCategory() string {
	switch t {
	case USER_EXTERNAL_AUTH_TYPE_OAUTH2_OIDC,
		USER_EXTERNAL_AUTH_TYPE_OAUTH2_NEXTCLOUD,
		USER_EXTERNAL_AUTH_TYPE_OAUTH2_GITEA,
		USER_EXTERNAL_AUTH_TYPE_OAUTH2_GITHUB:
		return USER_EXTERNAL_AUTH_TYPE_CATEOGRY_OAUTH2
	}
	return ""
}


func (t UserExternalAuthType) IsValid() bool {
	return t.GetCategory() != ""
}
