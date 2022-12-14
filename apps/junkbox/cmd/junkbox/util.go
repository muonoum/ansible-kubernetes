package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetLogFormat(format string) error {
	switch format {
	case "json":
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	case "console":
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr, PartsExclude: []string{"time"},
		})
	default:
		return fmt.Errorf("bad log format: %s", format)
	}

	return nil
}

func SetLogLevel(level string) error {
	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		return fmt.Errorf("bad log level: %s", level)
	}

	return nil
}
