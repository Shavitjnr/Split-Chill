package services

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/locales"
	"github.com/Shavitjnr/split-chill-ai/pkg/mail"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
	"github.com/Shavitjnr/split-chill-ai/pkg/templates"
)

const passwordResetUrlFormat = "%sdesktop#/resetpassword?token=%s"

type PasswordResetOTP struct {
	Email     string
	Code      string
	CreatedAt time.Time
	ExpiresAt time.Time
	Attempts  int
}

type ForgetPasswordService struct {
	ServiceUsingConfig
	ServiceUsingMailer
	otpMutex sync.RWMutex
	otpStore map[string]*PasswordResetOTP
}

var (
	ForgetPasswords = &ForgetPasswordService{
		ServiceUsingConfig: ServiceUsingConfig{
			container: settings.Container,
		},
		ServiceUsingMailer: ServiceUsingMailer{
			container: mail.Container,
		},
		otpStore: make(map[string]*PasswordResetOTP),
	}
)

func generateSecureOTP() string {
	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return fmt.Sprintf("%06d", mrand.Intn(900000)+100000)
	}
	return fmt.Sprintf("%06d", n.Int64()+100000)
}

func (s *ForgetPasswordService) GenerateAndSaveOTP(email string, ttl time.Duration) (string, error) {
	code := generateSecureOTP()
	s.otpMutex.Lock()
	defer s.otpMutex.Unlock()

	if s.otpStore == nil {
		s.otpStore = make(map[string]*PasswordResetOTP)
	}

	normalized := strings.ToLower(strings.TrimSpace(email))
	s.otpStore[normalized] = &PasswordResetOTP{
		Email:     normalized,
		Code:      code,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(ttl),
		Attempts:  0,
	}
	return code, nil
}

func (s *ForgetPasswordService) VerifyAndConsumeOTP(email string, code string) bool {
	s.otpMutex.Lock()
	defer s.otpMutex.Unlock()

	if s.otpStore == nil {
		return false
	}

	normalized := strings.ToLower(strings.TrimSpace(email))
	entry, exists := s.otpStore[normalized]
	if !exists {
		return false
	}

	if time.Now().After(entry.ExpiresAt) {
		delete(s.otpStore, normalized)
		return false
	}

	entry.Attempts++
	if entry.Attempts > 5 {
		delete(s.otpStore, normalized)
		return false
	}

	if strings.TrimSpace(code) != entry.Code {
		return false
	}

	delete(s.otpStore, normalized)
	return true
}

func (s *ForgetPasswordService) SendPasswordResetEmail(c core.Context, user *models.User, passwordResetToken string, otpCode string, backupLocale string) error {
	if !s.CurrentConfig().EnableSMTP {
		return errs.ErrSMTPServerNotEnabled
	}

	if otpCode == "" {
		otpCode, _ = s.GenerateAndSaveOTP(user.Email, 10*time.Minute)
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
			"OTPCode":             otpCode,
		},
	}

	var bodyBuffer bytes.Buffer
	err = tmpl.Execute(&bodyBuffer, templateParams)

	if err != nil {
		return err
	}

	message := &mail.MailMessage{
		To:      user.Email,
		Subject: fmt.Sprintf("[%s] %s - Your OTP is %s", localeTextItems.GlobalTextItems.AppName, forgetPasswordTextItems.Title, otpCode),
		Body:    bodyBuffer.String(),
	}

	err = s.SendMail(message)

	return err
}

