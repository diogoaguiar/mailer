package strategies

import (
	"sync"
	"time"

	"github.com/diogoaguiar/mailer/pkg/mailer/loggers"
)

// Sequential is a strategy that sends emails sequentially.
type Sequential struct {
	Interval int // Interval between emails, in seconds.
	Logger   loggers.Logger
	SendTo   func(subject string, body string, recipient string) error
}

// Send sends emails sequentially.
func (s *Sequential) Send(subject string, body string, recipients []string) error {
	var wg sync.WaitGroup

	for _, recipient := range recipients {
		wg.Add(1)
		s.Logger.Sending(recipient)

		go func(subject string, body string, recipient string) {
			defer wg.Done()

			if err := s.SendTo(subject, body, recipient); err != nil {
				s.Logger.Error(err)
				return
			}

			s.Logger.Success(recipient)
		}(subject, body, recipient)

		time.Sleep(time.Duration(s.Interval) * time.Second)
	}

	wg.Wait()

	return nil
}
