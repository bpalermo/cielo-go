package log

type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}
