package api

import (
	"fmt"
	"html"
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/mail"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)

type FeedbacksApi struct {
	ApiUsingConfig
}

var (
	Feedbacks = &FeedbacksApi{
		ApiUsingConfig: ApiUsingConfig{
			container: settings.Container,
		},
	}
)

func (a *FeedbacksApi) FeedbackSubmitHandler(c *core.WebContext) (any, *errs.Error) {
	var request models.FeedbackRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		log.Warnf(c, "[feedbacks.FeedbackSubmitHandler] parse request failed, because %s", err.Error())
		return nil, errs.ErrIncompleteOrIncorrectSubmission
	}

	if !a.CurrentConfig().EnableSMTP {
		log.Warnf(c, "[feedbacks.FeedbackSubmitHandler] SMTP is disabled on server")
		return nil, errs.ErrSMTPServerNotEnabled
	}

	// The administrator/owner destination email
	targetAdminEmail := "shavitdaloutra28@gmail.com"
	if a.CurrentConfig().SMTPConfig != nil && a.CurrentConfig().SMTPConfig.SMTPUser != "" {
		targetAdminEmail = a.CurrentConfig().SMTPConfig.SMTPUser
	}

	categoryDisplay := request.Category
	switch categoryDisplay {
	case "enterprise":
		categoryDisplay = "Enterprise Pricing"
	case "selfhost":
		categoryDisplay = "Self-Hosting Assistance"
	case "partnership":
		categoryDisplay = "Partnership & API"
	default:
		if categoryDisplay == "" || categoryDisplay == "general" {
			categoryDisplay = "General Question"
		}
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05 MST")
	safeName := html.EscapeString(request.FullName)
	safeEmail := html.EscapeString(request.Email)
	safeCategory := html.EscapeString(categoryDisplay)
	safeMessage := html.EscapeString(request.Message)

	adminEmailBody := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<style>
  body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif; background-color: #f8fafc; margin: 0; padding: 24px; }
  .card { max-width: 600px; margin: 0 auto; background: #ffffff; border: 1px solid #e2e8f0; border-radius: 16px; overflow: hidden; box-shadow: 0 4px 16px rgba(0,0,0,0.06); }
  .header { background: linear-gradient(135deg, #10b981 0%, #059669 100%); padding: 28px 32px; color: #ffffff; }
  .header h1 { margin: 0; font-size: 22px; font-weight: 800; }
  .header p { margin: 6px 0 0 0; font-size: 13px; opacity: 0.95; }
  .content { padding: 32px; }
  .table { width: 100%%; border-collapse: collapse; margin-bottom: 24px; }
  .table td { padding: 10px 0; border-bottom: 1px solid #f1f5f9; font-size: 13px; }
  .table .label { color: #64748b; font-weight: 600; width: 130px; }
  .table .val { color: #0f172a; font-weight: 700; }
  .badge { display: inline-block; padding: 4px 10px; background: #ecfdf5; color: #047857; border-radius: 9999px; font-size: 12px; font-weight: 700; text-transform: uppercase; }
  .msg-box { background: #f8fafc; border: 1px solid #e2e8f0; border-radius: 12px; padding: 20px; margin-bottom: 24px; }
  .msg-label { font-size: 12px; font-weight: 700; text-transform: uppercase; color: #64748b; margin-bottom: 8px; letter-spacing: 0.5px; }
  .msg-text { font-size: 14px; line-height: 1.6; color: #1e293b; white-space: pre-wrap; word-break: break-word; }
  .btn-wrap { text-align: center; margin-top: 24px; }
  .reply-btn { display: inline-block; padding: 12px 28px; background: #059669; color: #ffffff; text-decoration: none; border-radius: 8px; font-weight: 700; font-size: 14px; }
  .footer { background: #f8fafc; border-top: 1px solid #e2e8f0; padding: 16px 32px; text-align: center; font-size: 12px; color: #94a3b8; }
</style>
</head>
<body>
<div class="card">
  <div class="header">
    <h1>Spill Chill AI &mdash; New Feedback</h1>
    <p>Received from the landing page contact form</p>
  </div>
  <div class="content">
    <table class="table">
      <tr>
        <td class="label">Sender Name:</td>
        <td class="val">%s</td>
      </tr>
      <tr>
        <td class="label">Email Address:</td>
        <td class="val"><a href="mailto:%s" style="color: #059669; text-decoration: none;">%s</a></td>
      </tr>
      <tr>
        <td class="label">Inquiry Topic:</td>
        <td class="val"><span class="badge">%s</span></td>
      </tr>
      <tr>
        <td class="label">Submitted Time:</td>
        <td class="val" style="color: #64748b; font-weight: 500;">%s</td>
      </tr>
    </table>
    
    <div class="msg-box">
      <div class="msg-label">Message:</div>
      <div class="msg-text">%s</div>
    </div>
    
    <div class="btn-wrap">
      <a href="mailto:%s?subject=Re: Spill Chill Feedback (%s)" class="reply-btn">Reply Directly to %s</a>
    </div>
  </div>
  <div class="footer">
    Spill Chill AI Financial Platform &bull; Customer Support
  </div>
</div>
</body>
</html>`, safeName, safeEmail, safeEmail, safeCategory, nowStr, safeMessage, safeEmail, safeCategory, safeName)

	adminMsg := &mail.MailMessage{
		To:      targetAdminEmail,
		Subject: fmt.Sprintf("[Spill Chill Feedback] %s from %s", safeCategory, safeName),
		Body:    adminEmailBody,
	}

	sendErr := mail.Container.SendMail(adminMsg)
	if sendErr != nil {
		log.Errorf(c, "[feedbacks.FeedbackSubmitHandler] failed to send feedback email to admin %s: %s", targetAdminEmail, sendErr.Error())
		return nil, errs.ErrOperationFailed
	}

	log.Infof(c, "[feedbacks.FeedbackSubmitHandler] successfully sent feedback email to %s from %s <%s>", targetAdminEmail, safeName, safeEmail)

	// Send an automated receipt/acknowledgement to the user if different from admin
	if request.Email != "" && request.Email != targetAdminEmail {
		go func() {
			userReceiptBody := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<style>
  body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif; background-color: #f8fafc; margin: 0; padding: 24px; }
  .card { max-width: 600px; margin: 0 auto; background: #ffffff; border: 1px solid #e2e8f0; border-radius: 16px; overflow: hidden; box-shadow: 0 4px 16px rgba(0,0,0,0.06); }
  .header { background: linear-gradient(135deg, #10b981 0%, #059669 100%); padding: 28px 32px; color: #ffffff; }
  .header h1 { margin: 0; font-size: 20px; font-weight: 800; }
  .header p { margin: 6px 0 0 0; font-size: 13px; opacity: 0.95; }
  .content { padding: 32px; }
  .msg-box { background: #f8fafc; border: 1px solid #e2e8f0; border-radius: 12px; padding: 16px; margin: 20px 0; font-size: 13px; color: #475569; font-style: italic; }
  .footer { background: #f8fafc; border-top: 1px solid #e2e8f0; padding: 16px 32px; text-align: center; font-size: 12px; color: #94a3b8; }
</style>
</head>
<body>
<div class="card">
  <div class="header">
    <h1>We've received your message!</h1>
    <p>Thank you for contacting Spill Chill AI</p>
  </div>
  <div class="content">
    <p style="font-size: 14px; line-height: 1.6; color: #1e293b;">Hi <strong>%s</strong>,</p>
    <p style="font-size: 14px; line-height: 1.6; color: #334155;">Thank you for getting in touch with us regarding <strong>%s</strong>. Our team has received your message and will review it promptly. We usually reply within 24 hours.</p>
    <div class="msg-box">
      "%s"
    </div>
    <p style="font-size: 13px; color: #64748b; margin-top: 24px;">Best regards,<br><strong>The Spill Chill AI Team</strong></p>
  </div>
  <div class="footer">
    Spill Chill AI Financial Platform
  </div>
</div>
</body>
</html>`, safeName, safeCategory, safeMessage)

			receiptMsg := &mail.MailMessage{
				To:      request.Email,
				Subject: "We received your message — Spill Chill AI",
				Body:    userReceiptBody,
			}
			_ = mail.Container.SendMail(receiptMsg)
		}()
	}

	return &models.FeedbackResponse{
		Sent:    true,
		Message: "Thank you! Your feedback has been successfully submitted and delivered.",
	}, nil
}
