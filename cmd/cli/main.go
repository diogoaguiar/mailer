package main

import (
	"github.com/spf13/cobra"
)

var (
	strategy string
	sender   string
	file     bool
	to       []string
	subject  string
	body     string
)

var rootCmd = &cobra.Command{
	Use:   "mailer",
	Short: "A simple email sender",
	Long:  "A simple email sender",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&strategy, "strategy", "s", "sequential", "Strategy to use")
	rootCmd.PersistentFlags().StringVarP(&sender, "sender", "S", "smtp", "Sender service to use")
	rootCmd.PersistentFlags().BoolVarP(&file, "file", "f", false, "Read email body from file")
	rootCmd.PersistentFlags().StringArrayVarP(&to, "to", "t", []string{}, "Recipient email")

	rootCmd.Args = cobra.ExactArgs(2)
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		subject = args[0]
		body = args[1]
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	// print all arguments with a label
	println("strategy:", strategy)
	println("sender:", sender)
	println("file:", file)
	// print all recipients, enumerated
	for i, recipient := range to {
		println("recipient", i, ":", recipient)
	}
	println("subject:", subject)
	println("body:", body)
}
