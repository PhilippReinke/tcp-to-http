package logger

import (
	"log/slog"
)

type Level int

const (
	Unknown Level = iota
	Info
	Warn
	Error
	Debug
)

func (l Level) String() string {
	switch l {
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Debug:
		return "debug"
	default:
		return "unknown"
	}
}

func (l Level) slog() slog.Level {
	switch l {
	case Info:
		return slog.LevelInfo
	case Warn:
		return slog.LevelWarn
	case Error:
		return slog.LevelError
	case Debug:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}
