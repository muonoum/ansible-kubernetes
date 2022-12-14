package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/rs/zerolog/log"
)

func NewGren(optimize bool) api.Plugin {
	return api.Plugin{
		Name: "gren",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(
				api.OnResolveOptions{Filter: `\.gren$`},
				grenOnResolve,
			)

			build.OnLoad(
				api.OnLoadOptions{Filter: `.*`, Namespace: "gren"},
				grenOnLoad(optimize),
			)
		},
	}
}

func grenOnResolve(args api.OnResolveArgs) (api.OnResolveResult, error) {
	result := api.OnResolveResult{
		Path:      filepath.Join(args.ResolveDir, args.Path),
		Namespace: "gren",
	}

	return result, nil
}

func grenOnLoad(optimize bool) func(api.OnLoadArgs) (api.OnLoadResult, error) {
	return func(args api.OnLoadArgs) (api.OnLoadResult, error) {
		var result api.OnLoadResult

		if _, err := exec.LookPath("gren"); err != nil {
			return result, err
		}

		temp, err := os.CreateTemp("/tmp", "*.js")
		if err != nil {
			return result, err
		}
		defer os.Remove(temp.Name())

		wd, err := os.Getwd()
		if err != nil {
			return result, err
		}

		path, err := filepath.Rel(wd, args.Path)
		if err != nil {
			return result, err
		}

		buildCommand := []string{"gren", "make"}

		if optimize {
			buildCommand = append(buildCommand, "--optimize")
		}

		buildCommand = append(buildCommand, path, fmt.Sprintf("--output=%s", temp.Name()))
		cmd := exec.Command(buildCommand[0], buildCommand[1:]...)
		cmd.Stderr = os.Stderr

		log.Info().Str("plugin", "gren").Str("path", path).Msg("build")

		if err := cmd.Run(); err != nil {
			return result, err
		}

		compiled, err := os.ReadFile(temp.Name())
		if err != nil {
			return result, err
		}

		contents := string(compiled)
		result.Contents = &contents
		return result, nil
	}
}
