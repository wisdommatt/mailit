package mailit

// Mailer is the interface that wraps SendText and SendHTML
// methods.
type Mailer interface {
	SendText(dep TextDependencies) error
}

// SMTPConfig holds smtp configurations that a required to send the
// mail.
type SMTPConfig struct {
	Host, Username, Password string
	Port                     int
}

// mailer the object that implements the Mailer interface.
type mailer struct {
	smtpConfig SMTPConfig
}

// NewMailer returns a new mailer which implements the Mailer
// interface.
//
// NewMailer uses the informations in config to prefill the SMTP
// configurations for the mail.
func NewMailer(config SMTPConfig) Mailer {
	return &mailer{
		smtpConfig: config,
	}
}

// SendText sends a plain text emails.
// the emails can also be sent with attachments.
func (mailer *mailer) SendText(dep TextDependencies) (err error) {
	return
}
