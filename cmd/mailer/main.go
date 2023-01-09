package main

import (
	"os"
	"strconv"

	"github.com/diogoaguiar/mailer/internal/mailer"
	"github.com/diogoaguiar/mailer/internal/mailer/loggers"
	"github.com/diogoaguiar/mailer/internal/mailer/messages"
	"github.com/diogoaguiar/mailer/internal/mailer/senders"
	"github.com/diogoaguiar/mailer/internal/mailer/strategies"

	"github.com/joho/godotenv"
)

func main() {
	// Usage example
	godotenv.Load("config/.env")

	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	sender := senders.NewSMTPSender(
		host,
		port,
		username,
		password,
		messages.NewMessage(
			"Extended Warranty",
			"Hello, Sir.\n"+
				"We are contacting you to inform you that your warranty has expired.\n"+
				"Please contact us to extend it for a small fee of 1000€ per year of warranty.\n"+
				"Thank you for your attention and have a nice day :)",
		),
	)

	logger := loggers.NewConsoleLogger()

	interval, _ := strconv.Atoi(os.Getenv("STRATEGY_INTERVAL"))
	strategy := strategies.NewSequential(interval, logger)

	mailer := mailer.NewMailer(
		sender,
		strategy,
		logger,
		[]string{
			"user1@example.com",
			"user2@example.com",
			"user3@example.com",
		},
	)

	mailer.Send()
}
