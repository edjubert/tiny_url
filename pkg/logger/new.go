package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(levelStr string) *zerolog.Logger {
	var level zerolog.Level
	switch levelStr {
	case "trace":
		level = zerolog.TraceLevel
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	default:
		level = zerolog.InfoLevel
	}

	logger := zerolog.
		New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		Level(level).
		With().
		Timestamp().
		Caller().
		Logger()

	return &logger
}
