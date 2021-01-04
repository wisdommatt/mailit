package mailit

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
