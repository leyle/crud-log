package crudlog

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var defaultLogTarget = os.Stdout

func NewJsonLogger(level zerolog.Level) zerolog.Logger {
	zerolog.SetGlobalLevel(level)
	return zerolog.New(defaultLogTarget).With().Timestamp().Caller().Logger()
}

func NewConsoleLogger(level zerolog.Level) zerolog.Logger {
	zerolog.SetGlobalLevel(level)
	consoleWriter := zerolog.ConsoleWriter{
		Out:        defaultLogTarget,
		NoColor:    false,
		TimeFormat: time.RFC3339,
	}

	l := zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()
	return l
}
