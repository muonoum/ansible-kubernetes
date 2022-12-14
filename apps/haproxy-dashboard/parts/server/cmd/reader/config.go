package main

import "time"

type Config struct {
	CACert    string        `env:"READER_CA_CERT" required`
	SourceURL string        `env:"READER_SOURCE_URL" required`
	Timeout   time.Duration `env:"READER_TIMEOUT" default:"10s"`
	LogFormat string        `env:"READER_LOG_FORMAT" enum:"json,console" default:"console"`
	LogLevel  string        `env:"READER_LOG_LEVEL" enum:"debug,info,warn,error" default:"info"`

	Start StartCommand `cmd`
	Read  ReadCommand  `cmd`
}
