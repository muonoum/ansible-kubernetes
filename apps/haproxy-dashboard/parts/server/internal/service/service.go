package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Start(ctx context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() { <-interrupt; cancel() }()
	return ctx, cancel
}

func ConfigureLogger(format, level string) {
	switch format {
	case "json":
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	case "console":
		writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"}
		log.Logger = log.Output(writer)
	default:
		panic(format)
	}

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
		panic(level)
	}
}
