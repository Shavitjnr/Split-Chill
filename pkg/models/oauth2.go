package models


type OAuth2LoginRequest struct {
	Platform        string `form:"platform" binding:"required"`
	ClientSessionId string `form:"client_session_id" binding:"required"`
	Token           string `form:"token"`
}


type OAuth2CallbackRequest struct {
	State            string `form:"state"`
	Code             string `form:"code"`
	Error            string `form:"error"`
	ErrorDescription string `form:"error_description"`
}


type OAuth2CallbackLoginRequest struct {
	Password string `json:"password" binding:"omitempty,min=6,max=128"`
	Passcode string `json:"passcode" binding:"omitempty,notBlank,len=6"`
	Token    string `json:"token" binding:"omitempty"`
}
