package mailer

type Sender interface {
	SendTo(recipient string) error
}

type Strategy interface {
	Send(sender *Sender, recipients []string) error
}

type Logger interface {
	Success(recipient string)
	Error(err error)
}

type Mailer struct {
	sender     Sender
	strategy   Strategy
	logger     Logger
	recipients []string
}

func NewMailer(sender Sender, strategy Strategy, logger Logger, recipients []string) *Mailer {
	return &Mailer{
		sender:     sender,
		strategy:   strategy,
		logger:     logger,
		recipients: recipients,
	}
}

func (m *Mailer) Send() error {
	return m.strategy.Send(&m.sender, m.recipients)
}
