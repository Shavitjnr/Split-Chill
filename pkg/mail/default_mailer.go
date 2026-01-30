package mail

import (
	"crypto/tls"
	"net"

	"gopkg.in/mail.v2"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)


type DefaultMailer struct {
	dialer      *mail.Dialer
	fromAddress string
}


func NewDefaultMailer(smtpConfig *settings.SMTPConfig) (*DefaultMailer, error) {
	host, portStr, err := net.SplitHostPort(smtpConfig.SMTPHost)

	if err != nil {
		return nil, errs.ErrSMTPServerHostInvalid
	}

	port, err := utils.StringToInt(portStr)

	if err != nil {
		return nil, errs.ErrSMTPServerHostInvalid
	}

	dialer := mail.NewDialer(host, port, smtpConfig.SMTPUser, smtpConfig.SMTPPasswd)
	dialer.TLSConfig = &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: smtpConfig.SMTPSkipTLSVerify,
	}

	mailer := &DefaultMailer{
		dialer:      dialer,
		fromAddress: smtpConfig.FromAddress,
	}

	return mailer, nil
}


func (m *DefaultMailer) SendMail(message *MailMessage) error {
	if m.dialer == nil {
		return errs.ErrSMTPServerNotEnabled
	}

	mailMessage := mail.NewMessage()
	mailMessage.SetHeader("From", m.fromAddress)
	mailMessage.SetHeader("To", message.To)
	mailMessage.SetHeader("Subject", message.Subject)
	mailMessage.SetBody("text/html", message.Body)

	err := m.dialer.DialAndSend(mailMessage)

	return err
}
