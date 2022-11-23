package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	lib *logrus.Logger
}

func (l *Logger) Setup() {
	l.lib.SetFormatter(&logrus.TextFormatter{})
}

func (l *Logger) Debug(args ...interface{}) {
	l.lib.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.lib.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.lib.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.lib.Error(args...)
}

func New() *Logger {
	logger := &Logger{
		lib: logrus.New(),
	}

	logger.Setup()

	return logger
}
