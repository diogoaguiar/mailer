package mailer

import (
	"net/mail"

	"github.com/diogoaguiar/mailer/pkg/mailer/messages"
	"github.com/diogoaguiar/mailer/pkg/mailer/senders"
	"github.com/diogoaguiar/mailer/pkg/mailer/strategies"
)

// MailerConfig is the configuration for the mailer.
type MailerConfig struct {
	Message    messages.Message
	Sender     senders.Sender
	Strategy   strategies.Strategy
	Recipients []*mail.Address
}

// New
func New(config *MailerConfig) (*Mailer, error) {
	return &Mailer{
		message:    config.Message,
		sender:     config.Sender,
		strategy:   config.Strategy,
		recipients: config.Recipients,
	}, nil
}

type Mailer struct {
	message    messages.Message
	sender     senders.Sender
	strategy   strategies.Strategy
	recipients []*mail.Address
}

func (m *Mailer) Send() error {
	return m.strategy.Send(m.message.Subject(), m.message.Body(), m.recipients)
}
