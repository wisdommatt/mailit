package mailit

// TextDependencies is the struct that holds dependencies for
// a raw text email. e.g message body, sender/receiver email etc.
type TextDependencies struct {
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

	// Body is body of the email e.g Thanks for signing up
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
func (mailer *mailer) SendText(dep TextDependencies) (err error) {
	return
}
