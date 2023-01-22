package main

import (
	"fmt"
	"os"

	"github.com/diogoaguiar/mailer/internal/cli"
	"github.com/diogoaguiar/mailer/pkg/mailer"
	"github.com/diogoaguiar/mailer/pkg/mailer/configs"
	"github.com/diogoaguiar/mailer/pkg/mailer/messages"
	"github.com/diogoaguiar/mailer/pkg/mailer/senders"
	"github.com/diogoaguiar/mailer/pkg/mailer/strategies"
)

func main() {
	cli, err := cli.Parse()
	if err != nil {
		return
	}

	config := configs.Load()

	message := messages.NewMessage(cli.Subject, cli.Body)

	sender := getSender(cli.Sender, config)

	strategy := getStrategy(cli.Strategy, config, sender)

	mailerConfig := &mailer.MailerConfig{
		Message:    message,
		Sender:     sender,
		Strategy:   strategy,
		Recipients: cli.Recipients,
	}

	mailer, err := mailer.New(mailerConfig)
	if err != nil {
		fmt.Println("Something went wrong while creating the mailer:", err)
		os.Exit(1)
	}

	if err := mailer.Send(); err != nil {
		fmt.Println("Something went wrong while sending the emails:", err)
		os.Exit(2)
	}

	fmt.Println("Emails sent successfully!")
}

func getSender(sender string, config *configs.Config) senders.Sender {
	switch sender {
	case "smtp":
		return &senders.SMTPSender{
			Host:     config.Host,
			Port:     config.Port,
			Username: config.Username,
			Password: config.Password,
		}
	}

	return nil
}

func getStrategy(strategy string, config *configs.Config, sender senders.Sender) strategies.Strategy {
	switch strategy {
	case "sequential":
		return &strategies.Sequential{
			Interval: config.Interval,
			SendTo:   sender.SendTo,
		}
	}

	return nil
}
