package strategies

import (
	"log"
	"net/mail"
	"sync"
	"time"
)

// Sequential is a strategy that sends emails sequentially.
type Sequential struct {
	Interval int // Interval between emails, in seconds.
	SendTo   func(subject string, body string, recipient *mail.Address) error
}

// Send sends emails sequentially.
func (s *Sequential) Send(subject string, body string, recipients []*mail.Address) error {
	var wg sync.WaitGroup

	for _, recipient := range recipients {
		wg.Add(1)
		log.Default().Println("Sending to", recipient.String())

		go func(subject string, body string, recipient *mail.Address) {
			defer wg.Done()

			if err := s.SendTo(subject, body, recipient); err != nil {
				log.Default().Println("Error sending to", recipient.String(), err)
				return
			}

			log.Default().Println("Sent to", recipient.String())
		}(subject, body, recipient)

		time.Sleep(time.Duration(s.Interval) * time.Second)
	}

	wg.Wait()

	return nil
}
