package mailit

// TextDependencies is the struct that holds dependencies for
// a raw text email. e.g smtp host, sender/receiver email etc.
type TextDependencies struct {
	SenderEmail   string
	ReceiverEmail string
	Subject       string
	MessageBody   string
	Attachments   []string
}
