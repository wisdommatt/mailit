package mailit

import (
	"gopkg.in/gomail.v2"
)

// TextDependencies is the struct that holds dependencies for
// a plain text email. e.g message body, sender/receiver email etc.
type TextDependencies struct {
	// From is the email address of the sender
	// e.g no-reply@example.com
	From string

	// SenderName is the name fo the email sender.
	SenderName string

	// To is the receiver/receivers of the email
	// e.g []{"user@example.com", "user2@example.com"}.
	To []string

	// Subject is the subject of the email e.g Hello NewsLetter
	Subject string

	// Body is the body of the email e.g Thanks for signing up
	Body string

	// Attachments should hold the list of attachments to send with
	// the email if there is any.
	// 		[]{"assets/docs/welcome.pdf", "assets/images/admin.jpeg"}
	Attachments []string
}

// TextMailer is the interface that wraps SendText method.
type TextMailer interface {
	SendText(dep TextDependencies) error
}

// SendText sends a plain text emails.
// the emails can also be sent with attachments.
func (m *mailer) SendText(dep TextDependencies) error {
	mail := gomail.NewMessage()
	mail.SetAddressHeader("From", dep.From, dep.SenderName)
	mail.SetHeader("Subject", dep.Subject)
	mail.SetBody("text/plain", dep.Body)
	m.addRecipients(mail, dep.To)
	m.addAttachments(mail, dep.Attachments)
	mailDialer := gomail.NewDialer(m.smtp.Host, m.smtp.Port, m.smtp.Username, m.smtp.Password)
	err := mailDialer.DialAndSend(mail)
	if err != nil {
		return err
	}
	return nil
}
