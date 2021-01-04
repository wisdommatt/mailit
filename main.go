package mailit

// Mailer is the interface that wraps SendText and SendHTML
// methods.
type Mailer interface {
	SendText(dep TextDependencies)
}
