package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitGlobalLogger() {
	// Configure console writer for colorful logs
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    false, // Enable color in logs
	}

	// Set up the global logger to write only to console
	log.Logger = zerolog.New(consoleWriter).
		With().
		Timestamp().
		Caller().
		Logger()
}
