package crudlog

import (
	"github.com/rs/zerolog"
	"os"
	"strconv"
	"time"
)

var LogTarget = os.Stdout

func NewJsonLogger(level zerolog.Level) zerolog.Logger {
	zerolog.SetGlobalLevel(level)
	return zerolog.New(LogTarget).With().Timestamp().Caller().Logger()
}

func NewConsoleLogger(level zerolog.Level) zerolog.Logger {
	zerolog.SetGlobalLevel(level)

	// Lshortfile
	zerolog.CallerMarshalFunc = func(file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:        LogTarget,
		NoColor:    false,
		TimeFormat: time.RFC3339,
	}

	l := zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()
	return l
}
