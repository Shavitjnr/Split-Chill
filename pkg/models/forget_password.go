package models


type ForgetPasswordRequest struct {
	Email string `json:"email" binding:"required,notBlank,max=100,validEmail"`
}


type PasswordResetRequest struct {
	Email    string `json:"email" binding:"required,notBlank,max=100,validEmail"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}
