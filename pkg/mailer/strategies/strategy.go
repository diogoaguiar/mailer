package strategies

type Strategy interface {
	Send(subject string, body string, recipients []string) error
}
