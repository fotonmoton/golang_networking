package log

import (
	"go.uber.org/zap"
)

var logger *respLogger

type respLogger struct {
	Logger *zap.Logger
}

func getLogger() *respLogger {
	if logger == nil {
		l, _ := zap.NewDevelopment()
		defer l.Sync()

		logger = &respLogger{
			Logger: l,
		}
	}

	return logger
}

func (r respLogger) Warn(msg string) {
	r.Logger.Warn(msg)
}

func (r respLogger) Info(msg string) {
	r.Logger.Info(msg)
}

func (r respLogger) Error(msg string) {
	r.Logger.Error(msg)
}

func Warn(msg string) {
	l := getLogger()
	l.Warn(msg)
}

func Info(msg string) {
	l := getLogger()
	l.Info(msg)
}

func Error(msg string) {
	l := getLogger()
	l.Error(msg)
}
