package mail


type Mailer interface {
	SendMail(message *MailMessage) error
}
