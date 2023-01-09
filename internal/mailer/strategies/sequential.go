package strategies

import (
	"time"

	"github.com/diogoaguiar/mailer/internal/mailer"
)

// Sequential is a strategy that sends emails sequentially.
type Sequential struct {
	interval int // Interval between emails, in seconds.
	logger   mailer.Logger
}

// NewSequential creates a new Sequential strategy.
func NewSequential(interval int, logger mailer.Logger) *Sequential {
	return &Sequential{
		interval: interval,
		logger:   logger,
	}
}

// Send sends emails sequentially.
func (s *Sequential) Send(sender *mailer.Sender, recipients []string) error {
	for _, recipient := range recipients {
		if err := (*sender).SendTo(recipient); err != nil {
			s.logger.Error(err)
			continue
		}

		s.logger.Success(recipient)

		time.Sleep(time.Duration(s.interval) * time.Second)
	}

	return nil
}
