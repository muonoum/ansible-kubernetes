package main

import (
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Bundler struct {
	Optimize     bool          `help:"Optimized build where applicable."`
	Clear        bool          `help:"Clear console before rebuild."`
	Debug        bool          `help:"Show debug messages."`
	Debounce     time.Duration `help:"Debounce build." default:"250ms"`
	Load         []string      `help:"File loaders." placeholder:"ext"`
	Watch        string        `help:"Watch folder."`
	Output       string        `help:"Output folder." required`
	NotifyURL    string        `help:"URL to notify on successful build."`
	NotifyMethod string        `help:"Method to use when notifying." optional default:"PATCH"`
	Entrypoint   []string      `help:"Entrypoints to build." arg placeholder:"path"`
}

func main() {
	var cli Bundler
	kong.Parse(&cli)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stdout, PartsExclude: []string{"time"},
	})

	if cli.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Info().
		Bool("clear", cli.Clear).
		Dur("debounce", cli.Debounce).
		Bool("debug", cli.Debug).
		Strs("load", cli.Load).
		Str("notify-method", cli.NotifyMethod).
		Str("notify-url", cli.NotifyURL).
		Bool("optimize", cli.Optimize).
		Str("output", cli.Output).
		Str("watch", cli.Watch).
		Strs("entrypoints", cli.Entrypoint).
		Msg("bundle")

	if err := os.MkdirAll(cli.Output, 0750); err != nil {
		log.Fatal().Err(err).Msg("could not create output directory")
	}

	watcher, err := NewWatcher(cli.Watch)
	if err != nil {
		log.Fatal().Err(err).Msg("could not create watcher")
	}
	defer watcher.Close()

	var builds []*Build
	for _, path := range cli.Entrypoint {
		build := NewBuild(path, cli.Output, cli.Optimize, cli.Load)
		builds = append(builds, build)
	}

	if cli.Watch != "" {
		if cli.Clear {
			// fmt.Fprint(os.Stderr, "\033[H\033[2J")
		}

		maybeNotify(cli.NotifyMethod, cli.NotifyURL)

		watcher.Watch(cli.Debounce, cli.NotifyMethod, cli.NotifyURL, func(name string) {
			for _, build := range builds {
				if cli.Clear {
					// fmt.Fprint(os.Stderr, "\033[H\033[2J")
				}
				build.Rebuild()
			}
		})
	}
}
