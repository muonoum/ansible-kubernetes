package main

import (
	"github.com/alecthomas/kong"

	"junkbox/internal/cmds/node_hash"
	"junkbox/internal/cmds/split_yaml"
)

type CLI struct {
	LogFormat string `env:"LOG_FORMAT" enum:"json,console" default:"console"`
	LogLevel  string `env:"LOG_LEVEL" enum:"debug,info,warn,error" default:"info"`

	SplitYAML split_yaml.Command `cmd`
	NodeHash  node_hash.Command  `cmd`
}

func main() {
	var err error
	var cli CLI
	ctx := kong.Parse(&cli)

	err = SetLogFormat(cli.LogFormat)
	ctx.FatalIfErrorf(err)

	err = SetLogLevel(cli.LogLevel)
	ctx.FatalIfErrorf(err)

	err = ctx.Run(cli)
	ctx.FatalIfErrorf(err)
}
