package service

type Config struct {
	Address   string `env:"ADDRESS" required`
	LogFormat string `env:"LOG_FORMAT" enum:"json,console" default:"console"`
	LogLevel  string `env:"LOG_LEVEL" enum:"debug,info,warn,error" default:"info"`
}
