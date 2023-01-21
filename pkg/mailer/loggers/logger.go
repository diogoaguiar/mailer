package loggers

type Logger interface {
	Sending(recipient string)
	Success(recipient string)
	Error(err error)
}
