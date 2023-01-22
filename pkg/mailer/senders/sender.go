package senders

import "net/mail"

type Sender interface {
	SendTo(subject string, body string, recipient *mail.Address) error
}
