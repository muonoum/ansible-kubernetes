package main

import (
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	Address   string        `env:"DASHBOARD_ADDRESS" required`
	Source    string        `env:"DASHBOARD_SOURCE" `
	Timeout   time.Duration `env:"READER_TIMEOUT" default:"10s"`
	LogFormat string        `env:"DASHBOARD_LOG_FORMAT" enum:"json,console" default:"console"`
	LogLevel  string        `env:"DASHBOARD_LOG_LEVEL" enum:"debug,info,warn,error" default:"info"`
}

func (config Config) MarshalZerologObject(event *zerolog.Event) {
	event.Str("ADDRESS", config.Address).
		Str("SOURCE", config.Source).
		Dur("TIMEOUT", config.Timeout).
		Str("LOG_FORMAT", config.LogFormat).
		Str("LOG_LEVEL", config.LogLevel)
}
