package senders

type Sender interface {
	SendTo(subject string, body string, recipient string) error
}
