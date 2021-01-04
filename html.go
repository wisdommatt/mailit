package mailit

import (
	"bytes"
	"text/template"

	"gopkg.in/gomail.v2"
)

// HTMLDependencies is the struct that holds dependencies for
// a html email. e.g template, sender/receiver email etc.
type HTMLDependencies struct {
	// From is the email address of the sender
	// e.g no-reply@example.com
	From string

	// To is the receiver/receivers of the email
	// e.g user@example.com.
	// if the receivers are more one seperate them with
	// a comma and space e.g user1@example.com, user2@example.com
	To string

	// Subject is the subject of the email e.g Hello NewsLetter
	Subject string

	// Template is the directory location of the HTML template
	// e.g templates/welcome.html
	Template string

	// TemplateData is the data/content you want to pass to the HTML
	// template e.g User Info, Order Details etc
	TemplateData interface{}

	// Attachments should hold the list of attachments to send with
	// the email if there is any.
	// 		[]{"assets/docs/welcome.pdf", "assets/images/admin.jpeg"}
	Attachments []string
}

// HTMLMailer is the interface that wraps SendHTML method.
type HTMLMailer interface {
	SendHTML(dep HTMLDependencies) error
}

// SendHTML sends a email using a HTML template.
// It uses the informations in dep to forward the email
// and it also supports attachments.
func (m *mailer) SendHTML(dep HTMLDependencies) error {
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
	mail.SetHeader("From", dep.From)
	mail.SetHeader("To", dep.To)
	mail.SetHeader("Subject", dep.Subject)
	mail.SetBody("text/html", tempString)
	mailDialer := gomail.NewDialer(m.smtp.Host, m.smtp.Port, m.smtp.Username, m.smtp.Password)
	err = mailDialer.DialAndSend(mail)
	if err != nil {
		return err
	}
	return nil
}
