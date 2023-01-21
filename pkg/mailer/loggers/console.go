package loggers

import (
	"fmt"
	"os"
)

// ConsoleLogger is a logger that prints to the console.
type ConsoleLogger struct {
}

func (l ConsoleLogger) Sending(recipient string) {
	fmt.Println("Sending email to " + recipient)
}

func (l ConsoleLogger) Success(recipient string) {
	fmt.Println("Email sent to " + recipient)
}

func (l ConsoleLogger) Error(err error) {
	fmt.Fprintln(os.Stderr, err)
}
