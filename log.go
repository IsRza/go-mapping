package mapping

import (
	"context"
	"log"
)

var logger Logger
var loggerProvider func(context.Context) Logger

type (
	Logger interface {
		Error(args ...any)
		Errorf(format string, args ...any)

		Warn(args ...any)
		Warnf(format string, args ...any)
	}

	defaultLogger struct{}
)

func (dl defaultLogger) Error(args ...any) {
	log.Print(args...)
}

func (dl defaultLogger) Errorf(format string, args ...any) {
	log.Printf(format, args...)
}

func (dl defaultLogger) Warn(args ...any) {
	log.Print(args...)
}

func (dl defaultLogger) Warnf(format string, args ...any) {
	log.Printf(format, args...)
}

func init() {
	logger = defaultLogger{}
}

func SetLogger(l Logger) {
	logger = l
}

func SetLoggerProvider(lp func(context.Context) Logger) {
	loggerProvider = lp
}

func getLogger(ctx context.Context) Logger {
	if loggerProvider != nil && ctx != nil {
		return loggerProvider(ctx)
	}
	return logger
}
