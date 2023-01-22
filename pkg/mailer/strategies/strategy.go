package strategies

import "net/mail"

type Strategy interface {
	Send(subject string, body string, recipients []*mail.Address) error
}
