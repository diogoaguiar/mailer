package cli

import (
	"fmt"
	"log"
	"net/mail"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Cli struct {
	Strategy   string
	Sender     string
	Recipients []*mail.Address
	Subject    string
	Body       string
}

func (c *Cli) String() string {
	return fmt.Sprintf(
		"strategy: %s,\n"+
			"sender: %s,\n"+
			"to: %s,\n"+
			"subject: %s,\n"+
			"body: %s",
		c.Strategy,
		c.Sender,
		c.Recipients,
		c.Subject,
		c.Body,
	)
}

var (
	cli     = &Cli{}
	rootCmd = &cobra.Command{
		Use:   "mailer",
		Short: "A simple email sender",
	}
	body       string
	recipients string
)

func init() {
	rootCmd.Flags().StringVarP(&cli.Subject, "subject", "s", "", "Subject of the email")
	rootCmd.Flags().StringVarP(&body, "body", "b", "", "Body of the email")
	rootCmd.Flags().StringVarP(&recipients, "to", "t", "", "Recipient emails")
	rootCmd.Flags().StringVarP(&cli.Strategy, "strategy", "", "sequential", "Strategy to use")
	rootCmd.Flags().StringVarP(&cli.Sender, "sender", "", "smtp", "Sender service to use")
	rootCmd.MarkFlagRequired("subject")
	rootCmd.MarkFlagRequired("body")
	rootCmd.MarkFlagRequired("to")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		cli.Body = parseBody(body)
		cli.Recipients = parseRecipients(recipients)
	}
}

func Parse() (*Cli, error) {
	if err := rootCmd.Execute(); err != nil {
		return nil, err
	}

	return cli, nil
}

func parseBody(body string) string {
	if isAFile(body) {
		bytes, err := os.ReadFile(body)
		if err != nil {
			log.Fatalln("failed to read body file: ", err)
		}

		body = string(bytes)
	}

	return body
}

func parseRecipients(recipients string) []*mail.Address {
	if isAFile(recipients) {
		bytes, err := os.ReadFile(recipients)
		if err != nil {
			log.Fatalln("failed to read recipients file: ", err)
		}

		recipients = string(bytes)
	}

	recipients = strings.Map(func(r rune) rune {
		switch r {
		case ';', '\n':
			r = ','
		}

		return r
	}, recipients)

	addresses, err := mail.ParseAddressList(recipients)
	if err != nil {
		log.Fatalln("failed to parse recipients: ", err)
	}

	return addresses
}

func isAFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
