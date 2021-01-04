package mailit

import (
	"log"

	"gopkg.in/gomail.v2"
)

// Mailer is the interface that wraps SendText and SendHTML
// methods.
type Mailer interface {
	TextMailer
	HTMLMailer
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

// addAttachments adds attachments to an email before it is sent
// out.
func (m *mailer) addAttachments(mail *gomail.Message, attachments []string) {
	for _, attachment := range attachments {
		log.Println("attaching ", attachment)
		mail.Attach(attachment)
	}
}
