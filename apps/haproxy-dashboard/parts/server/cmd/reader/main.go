package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var config Config
	ctx := kong.Parse(&config)
	err := ctx.Validate()
	ctx.FatalIfErrorf(err)

	switch config.LogFormat {
	case "json":
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	case "console":
		writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"}
		log.Logger = log.Output(writer)
	default:
		panic(config.LogFormat)
	}

	switch config.LogLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		panic(config.LogLevel)
	}

	err = ctx.Run(config)
	ctx.FatalIfErrorf(err)
}
