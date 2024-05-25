package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func New(level string) *zerolog.Logger {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	logger := zerolog.New(os.Stderr).Level(logLevel).With().Timestamp().Logger()

	return &logger
}
