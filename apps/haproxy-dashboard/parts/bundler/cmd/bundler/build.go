package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"bundler/internal/plugin"
)

type Build struct {
	entry  string
	output string
	result api.BuildResult
}

func NewBuild(entry, output string, optimize bool, files []string) *Build {
	var options api.BuildOptions

	options.EntryPoints = []string{entry}
	options.Bundle = true
	options.Outdir = output
	options.EntryNames = "[dir]/[name]"
	options.AssetNames = "[dir]/[name]"
	options.Write = true
	options.MinifyWhitespace = optimize
	options.MinifyIdentifiers = optimize
	options.MinifySyntax = optimize
	// options.Target = api.ES5
	options.Incremental = true
	options.Plugins = []api.Plugin{
		plugin.NewGren(optimize),
	}

	options.Loader = make(map[string]api.Loader)
	for _, ext := range files {
		ext = fmt.Sprintf(".%s", ext)
		log.Debug().Str("ext", ext).Msg("file loader")
		options.Loader[ext] = api.LoaderFile
	}

	build := new(Build)
	build.entry = entry
	build.output = output
	build.result = api.Build(options)
	build.handle()

	return build
}

func (build *Build) Rebuild() {
	build.result = build.result.Rebuild()
	build.handle()
}

func (build *Build) write() error {
	var group errgroup.Group

	for _, file := range build.result.OutputFiles {
		file := file

		group.Go(func() error {
			path := filepath.Join(build.output, filepath.Base(file.Path))
			log.Debug().Str("source", build.entry).Str("path", path).Msg("output")
			return os.WriteFile(path, file.Contents, 0640)
		})
	}

	return group.Wait()
}

// TODO: Return error
func (build *Build) handle() {
	if len(build.result.Errors) > 0 {
		for _, err := range build.result.Errors {
			log.Error().Str("source", build.entry).Str("error", err.Text).Msg("build error")
		}

		log.Fatal().Msg("build failed")

		return
	}

	// if err := build.write(); err != nil {
	// 	log.Error().Err(err).Msg("write error")
	// }

	log.Info().Str("source", build.entry).Msg("build ok")
}
