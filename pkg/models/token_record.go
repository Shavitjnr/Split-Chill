package models

import "github.com/Shavitjnr/split-chill-ai/pkg/core"


const TokenMaxUserAgentLength = 255


type TokenRecord struct {
	Uid              int64          `xorm:"PK INDEX(IDX_token_record_uid_type_expired_time) INDEX(IDX_token_record_expired_time)"`
	UserTokenId      int64          `xorm:"PK"`
	TokenType        core.TokenType `xorm:"INDEX(IDX_token_record_uid_type_expired_time) TINYINT NOT NULL"`
	Secret           string         `xorm:"VARCHAR(10) NOT NULL"`
	UserAgent        string         `xorm:"VARCHAR(255)"`
	Context          string         `xorm:"BLOB"`
	CreatedUnixTime  int64          `xorm:"PK"`
	ExpiredUnixTime  int64          `xorm:"INDEX(IDX_token_record_uid_type_expired_time) INDEX(IDX_token_record_expired_time)"`
	LastSeenUnixTime int64
}


type OAuth2CallbackTokenContext struct {
	ExternalAuthType core.UserExternalAuthType `json:"externalAuthType"`
	ExternalUsername string                    `json:"externalUsername"`
	ExternalEmail    string                    `json:"externalEmail"`
}


type TokenGenerateAPIRequest struct {
	ExpiredInSeconds int64  `json:"expiresInSeconds" binding:"omitempty,min=0,max=4294967295"`
	Password         string `json:"password" binding:"omitempty,min=6,max=128"`
}


type TokenGenerateMCPRequest struct {
	ExpiredInSeconds int64  `json:"expiresInSeconds" binding:"omitempty,min=0,max=4294967295"`
	Password         string `json:"password" binding:"omitempty,min=6,max=128"`
}


type TokenRevokeRequest struct {
	TokenId string `json:"tokenId" binding:"required,notBlank"`
}


type TokenGenerateAPIResponse struct {
	Token      string `json:"token"`
	APIBaseUrl string `json:"apiBaseUrl"`
}


type TokenGenerateMCPResponse struct {
	Token  string `json:"token"`
	MCPUrl string `json:"mcpUrl"`
}


type TokenRefreshResponse struct {
	NewToken                 string                        `json:"newToken,omitempty"`
	OldTokenId               string                        `json:"oldTokenId,omitempty"`
	User                     *UserBasicInfo                `json:"user"`
	ApplicationCloudSettings *ApplicationCloudSettingSlice `json:"applicationCloudSettings,omitempty"`
	NotificationContent      string                        `json:"notificationContent,omitempty"`
}


type TokenInfoResponse struct {
	TokenId   string         `json:"tokenId"`
	TokenType core.TokenType `json:"tokenType"`
	UserAgent string         `json:"userAgent"`
	LastSeen  int64          `json:"lastSeen"`
	IsCurrent bool           `json:"isCurrent"`
}


type TokenInfoResponseSlice []*TokenInfoResponse


func (a TokenInfoResponseSlice) Len() int {
	return len(a)
}


func (a TokenInfoResponseSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}


func (a TokenInfoResponseSlice) Less(i, j int) bool {
	return a[i].LastSeen > a[j].LastSeen
}
