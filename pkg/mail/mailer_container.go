package mail

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type MailerContainer struct {
	current Mailer
}


var (
	Container = &MailerContainer{}
)


func InitializeMailer(config *settings.Config) error {
	if !config.EnableSMTP {
		Container.current = nil
		return nil
	}

	mailer, err := NewDefaultMailer(config.SMTPConfig)

	if err != nil {
		return err
	}

	Container.current = mailer
	return nil
}


func (m *MailerContainer) SendMail(message *MailMessage) error {
	if m.current == nil {
		return errs.ErrSMTPServerNotEnabled
	}

	return m.current.SendMail(message)
}
