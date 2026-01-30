package services

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/locales"
	"github.com/Shavitjnr/split-chill-ai/pkg/mail"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
	"github.com/Shavitjnr/split-chill-ai/pkg/templates"
)

const passwordResetUrlFormat = "%sdesktop#/resetpassword?token=%s"


type ForgetPasswordService struct {
	ServiceUsingConfig
	ServiceUsingMailer
}


var (
	ForgetPasswords = &ForgetPasswordService{
		ServiceUsingConfig: ServiceUsingConfig{
			container: settings.Container,
		},
		ServiceUsingMailer: ServiceUsingMailer{
			container: mail.Container,
		},
	}
)


func (s *ForgetPasswordService) SendPasswordResetEmail(c core.Context, user *models.User, passwordResetToken string, backupLocale string) error {
	if !s.CurrentConfig().EnableSMTP {
		return errs.ErrSMTPServerNotEnabled
	}

	locale := user.Language

	if locale == "" {
		locale = backupLocale
	}

	localeTextItems := locales.GetLocaleTextItems(locale)
	forgetPasswordTextItems := localeTextItems.ForgetPasswordMailTextItems

	expireTimeInMinutes := s.CurrentConfig().PasswordResetTokenExpiredTimeDuration.Minutes()
	passwordResetUrl := fmt.Sprintf(passwordResetUrlFormat, s.CurrentConfig().RootUrl, url.QueryEscape(passwordResetToken))

	tmpl, err := templates.GetTemplate(templates.TEMPLATE_PASSWORD_RESET)

	if err != nil {
		return err
	}

	templateParams := map[string]any{
		"AppName": localeTextItems.GlobalTextItems.AppName,
		"ForgetPasswordMail": map[string]any{
			"Title":               forgetPasswordTextItems.Title,
			"Salutation":          fmt.Sprintf(forgetPasswordTextItems.SalutationFormat, user.Nickname),
			"DescriptionAboveBtn": forgetPasswordTextItems.DescriptionAboveBtn,
			"ResetPasswordUrl":    passwordResetUrl,
			"ResetPassword":       forgetPasswordTextItems.ResetPassword,
			"DescriptionBelowBtn": fmt.Sprintf(forgetPasswordTextItems.DescriptionBelowBtnFormat, expireTimeInMinutes),
		},
	}

	var bodyBuffer bytes.Buffer
	err = tmpl.Execute(&bodyBuffer, templateParams)

	if err != nil {
		return err
	}

	message := &mail.MailMessage{
		To:      user.Email,
		Subject: forgetPasswordTextItems.Title,
		Body:    bodyBuffer.String(),
	}

	err = s.SendMail(message)

	return err
}
