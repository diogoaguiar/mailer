package loggers

import (
	"fmt"
	"os"
)

// ConsoleLogger is a logger that prints to the console.
type ConsoleLogger struct {
}

// NewConsoleLogger creates a new ConsoleLogger.
func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

// Success prints a success message to the console.
func (l ConsoleLogger) Success(recipient string) {
	fmt.Println("Email sent to " + recipient)
}

// Error prints an error message to the console.
func (l ConsoleLogger) Error(err error) {
	fmt.Fprintln(os.Stderr, err)
}
