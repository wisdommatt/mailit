package mailit

import (
	"gopkg.in/gomail.v2"
)

// Mailer is the interface that wraps TextMailer and
// TemplateMailer.
type Mailer interface {
	TextMailer
	TemplateMailer
}

// SMTPConfig holds smtp configurations that a required to send the
// mail.
type SMTPConfig struct {
	Host, Username, Password string
	Port                     int
}

// mailer the object that implements the Mailer interface.
type mailer struct {
	smtp SMTPConfig
}

// NewMailer returns a new mailer which implements the Mailer
// interface.
//
// NewMailer uses the informations in config to prefill the SMTP
// configurations for the mail.
func NewMailer(config SMTPConfig) Mailer {
	return &mailer{
		smtp: config,
	}
}

// addRecipients formats a slice of string containing recipients email
// and then adds it to the email.
func (m *mailer) addRecipients(mail *gomail.Message, recipients []string) {
	addresses := make([]string, len(recipients))
	for i, recipient := range recipients {
		addresses[i] = mail.FormatAddress(recipient, "")
	}
	mail.SetHeader("To", addresses...)
}

// addAttachments adds attachments to an email before it is sent
// out.
func (m *mailer) addAttachments(mail *gomail.Message, attachments []string) {
	for _, attachment := range attachments {
		mail.Attach(attachment)
	}
}
