package loggers

type Logger interface {
	Success(recipient string)
	Error(err error)
}
