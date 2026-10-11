package models


type FeedbackRequest struct {
	FullName string `json:"fullName" binding:"required,notBlank,min=2,max=100"`
	Email    string `json:"email" binding:"required,notBlank,max=100,validEmail"`
	Category string `json:"category"`
	Message  string `json:"message" binding:"required,notBlank,min=5,max=5000"`
}


type FeedbackResponse struct {
	Sent    bool   `json:"sent"`
	Message string `json:"message"`
}
