package logger

import (
	"log/slog"
	"net"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func New(level Level) *Logger {
	opts := &slog.HandlerOptions{
		Level: level.slog(),
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	return &Logger{logger: logger}
}

func (l *Logger) Debug(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) WithField(field string, val any) *Logger {
	attr := slog.Attr{
		Key:   field,
		Value: slog.AnyValue(val),
	}
	return &Logger{logger: l.logger.With(attr)}
}

func (l *Logger) WithError(err error) *Logger {
	return l.WithField("err", err.Error())
}

func (l *Logger) WithConnection(conn net.Conn) *Logger {
	return l.WithField("remote_addr", conn.RemoteAddr())
}
