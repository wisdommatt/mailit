package mailit

import (
	"bytes"
	"text/template"

	"gopkg.in/gomail.v2"
)

// TemplateDependencies is the struct that holds dependencies for
// a template based email. e.g template directory, sender/receiver email etc.
type TemplateDependencies struct {
	// From is the email address of the sender
	// e.g no-reply@example.com
	From string

	// To is the receiver/receivers of the email
	// e.g []{"user@example.com", "user2@example.com"}.
	To []string

	// SenderName is the name fo the email sender.
	SenderName string

	// Subject is the subject of the email e.g Hello NewsLetter
	Subject string

	// ContentType is the content type of the template e.g
	// text/plain or text/html etc
	ContentType string

	// Template is the directory location of the template
	// e.g templates/welcome.html, templates/hello.txt
	Template string

	// TemplateData is the data/content you want to pass to the
	// template e.g User Info, Order Details etc
	TemplateData interface{}

	// Attachments should hold the list of attachments to send with
	// the email if there is any.
	// 		[]{"assets/docs/welcome.pdf", "assets/images/admin.jpeg"}
	Attachments []string
}

// TemplateMailer is the interface that wraps SendTemplate method.
type TemplateMailer interface {
	SendTemplate(dep TemplateDependencies) error
}

// SendTemplate sends a email using a Text/HTML template.
// It uses the informations in dep to forward the email
// and it also supports attachments.
func (m *mailer) SendTemplate(dep TemplateDependencies) error {
	temp, err := template.ParseFiles(dep.Template)
	if err != nil {
		return err
	}
	var templateBuffer bytes.Buffer
	err = temp.Execute(&templateBuffer, dep.TemplateData)
	if err != nil {
		return err
	}
	tempString := templateBuffer.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", dep.From, dep.SenderName)
	mail.SetHeader("Subject", dep.Subject)
	mail.SetBody(dep.ContentType, tempString)
	m.addRecipients(mail, dep.To)
	m.addAttachments(mail, dep.Attachments)
	mailDialer := gomail.NewDialer(m.smtp.Host, m.smtp.Port, m.smtp.Username, m.smtp.Password)
	err = mailDialer.DialAndSend(mail)
	if err != nil {
		return err
	}
	return nil
}
