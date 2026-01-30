package models

import "github.com/Shavitjnr/split-chill-ai/pkg/core"


type UserExternalAuth struct {
	Uid              int64                     `xorm:"PK"`
	ExternalAuthType core.UserExternalAuthType `xorm:"VARCHAR(32) PK UNIQUE(uqe_userexternalauth_authtype_username) UNIQUE(uqe_userexternalauth_authtype_email)"`
	ExternalUsername string                    `xorm:"VARCHAR(32) UNIQUE(uqe_userexternalauth_authtype_username) NOT NULL"`
	ExternalEmail    string                    `xorm:"VARCHAR(100) UNIQUE(uqe_userexternalauth_authtype_email) NOT NULL"`
	CreatedUnixTime  int64
}


type UserExternalAuthUnlinkRequest struct {
	ExternalAuthType string `json:"externalAuthType" binding:"required,notBlank"`
	Password         string `json:"password" binding:"required,min=6,max=128"`
}


type UserExternalAuthInfoResponse struct {
	ExternalAuthCategory string                    `json:"externalAuthCategory"`
	ExternalAuthType     core.UserExternalAuthType `json:"externalAuthType"`
	Linked               bool                      `json:"linked"`
	ExternalUsername     string                    `json:"externalUsername,omitempty"`
	CreatedAt            int64                     `json:"createdAt,omitempty"`
}


func (a *UserExternalAuth) ToUserExternalAuthInfoResponse() *UserExternalAuthInfoResponse {
	return &UserExternalAuthInfoResponse{
		ExternalAuthCategory: a.ExternalAuthType.GetCategory(),
		ExternalAuthType:     a.ExternalAuthType,
		Linked:               true,
		ExternalUsername:     a.ExternalUsername,
		CreatedAt:            a.CreatedUnixTime,
	}
}


type UserExternalAuthInfoResponsesSlice []*UserExternalAuthInfoResponse


func (a UserExternalAuthInfoResponsesSlice) Len() int {
	return len(a)
}


func (a UserExternalAuthInfoResponsesSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}


func (a UserExternalAuthInfoResponsesSlice) Less(i, j int) bool {
	if a[i].Linked && !a[j].Linked {
		return true
	} else if !a[i].Linked && a[j].Linked {
		return false
	} else if !a[i].Linked && !a[j].Linked {
		return a[i].ExternalAuthType < a[j].ExternalAuthType
	}

	return a[i].CreatedAt > a[j].CreatedAt
}
