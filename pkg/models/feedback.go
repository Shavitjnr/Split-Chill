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


type Feedback struct {
	FeedbackId      int64  `xorm:"PK autoincr 'feedback_id'"`
	FullName        string `xorm:"VARCHAR(100) NOT NULL 'full_name'"`
	Email           string `xorm:"VARCHAR(100) NOT NULL 'email'"`
	Category        string `xorm:"VARCHAR(50) NOT NULL 'category'"`
	Message         string `xorm:"TEXT NOT NULL 'message'"`
	ClientIp        string `xorm:"VARCHAR(50) 'client_ip'"`
	Status          string `xorm:"VARCHAR(20) NOT NULL 'status'"`
	CreatedUnixTime int64  `xorm:"NOT NULL 'created_unix_time'"`
}
