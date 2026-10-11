package models


type ForgetPasswordRequest struct {
	Email string `json:"email" binding:"required,notBlank,max=100,validEmail"`
}


type PasswordResetRequest struct {
	Email    string `json:"email" binding:"required,notBlank,max=100,validEmail"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type PasswordResetByOTPRequest struct {
	Email    string `json:"email" binding:"required,notBlank,max=100,validEmail"`
	OTP      string `json:"otp" binding:"required,min=4,max=10"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

