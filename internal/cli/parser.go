package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Cli struct {
	Strategy   string
	Sender     string
	File       bool
	Recipients []string
	Subject    string
	Body       string
}

func (c *Cli) String() string {
	return fmt.Sprintf(
		"strategy: %s,\n"+
			"sender: %s,\n"+
			"file: %t,\n"+
			"to: %s,\n"+
			"subject: %s,\n"+
			"body: %s",
		c.Strategy,
		c.Sender,
		c.File,
		c.Recipients,
		c.Subject,
		c.Body,
	)
}

var (
	cli     = &Cli{}
	rootCmd = &cobra.Command{
		Use:   "mailer [flags] subject body",
		Short: "A simple email sender",
	}
)

func init() {
	rootCmd.Flags().StringVarP(&cli.Strategy, "strategy", "s", "sequential", "Strategy to use")
	rootCmd.Flags().StringVarP(&cli.Sender, "sender", "S", "smtp", "Sender service to use")
	rootCmd.Flags().BoolVarP(&cli.File, "file", "f", false, "Read email body from file")
	rootCmd.Flags().StringArrayVarP(&cli.Recipients, "to", "t", []string{}, "Recipient email")
	rootCmd.MarkFlagRequired("to")

	rootCmd.Args = cobra.ExactArgs(2)
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		cli.Subject = args[0]
		cli.Body = args[1]
	}
}

func Parse() (*Cli, error) {
	if err := rootCmd.Execute(); err != nil {
		return nil, err
	}

	return cli, nil
}
