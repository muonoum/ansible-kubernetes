package split_yaml

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	yaml "sigs.k8s.io/yaml"
)

type Command struct {
	Inputs    []string `arg`
	Default   string
	Kinds     map[string]string `name:"kind"`
	Overwrite bool
	DryRun    bool
}

func (cmd Command) Run() error {
	roots := make(map[string][]string)

	for _, input := range cmd.Inputs {
		log.Debug().Str("path", input).Msg("read")

		r, err := reader(input)
		if err != nil {
			return err
		}

		for yr := yamlutil.NewYAMLReader(bufio.NewReader(r)); ; {
			data, err := yr.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}

			var doc Document
			if err := yaml.Unmarshal(data, &doc); err != nil {
				return err
			}

			if doc.APIVersion == "" || doc.Kind == "" {
				continue
			}

			nameTemplate := filepath.Base(cmd.Default)
			relativeOutput := filepath.Dir(cmd.Default)
			if kind, ok := cmd.Kinds[doc.Kind]; ok {
				nameTemplate = filepath.Base(kind)
				relativeOutput = filepath.Dir(kind)
			}

			parsedTemplate, err := template.New("doc").Parse(nameTemplate)
			if err != nil {
				return err
			}

			var renderedFilename bytes.Buffer
			if err := parsedTemplate.Execute(&renderedFilename, doc); err != nil {
				return err
			}

			outputPath := filepath.Join(relativeOutput, renderedFilename.String())
			outputDir := filepath.Dir(outputPath)
			roots[outputDir] = append(roots[outputDir], filepath.Base(outputPath))

			if _, err := os.Stat(outputPath); err == nil && !cmd.Overwrite {
				log.Warn().Str("path", outputPath).Msg("skipping existing file")
				continue
			}

			log.Debug().Str("path", outputPath).Msg("output")

			if err := os.MkdirAll(outputDir, 0750); err != nil {
				return err
			}

			f, err := os.Create(outputPath)
			if err != nil {
				return err
			} else if _, err := f.Write(data); err != nil {
				return err
			}
			f.Close()
		}
	}

	for dir, resources := range roots {
		if err := writeKustomization(dir, resources); err != nil {
			return err
		}
	}

	return nil
}

func reader(path string) (io.Reader, error) {
	if path == "-" {
		return os.Stdin, nil
	}

	return os.Open(path)
}

func writeKustomization(dir string, resources []string) error {
	path := filepath.Join(dir, "kustomization.yaml")

	log.Info().
		Str("path", path).
		Int("resources", len(resources)).
		Msg("kustomization")

	lines := []string{
		"apiVersion: kustomize.config.k8s.io/v1beta1",
		"kind: Kustomization",
		"",
		"resources:",
	}

	for _, resource := range resources {
		lines = append(lines, fmt.Sprintf("  - %s", resource))
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := fmt.Fprintln(f, strings.Join(lines, "\n")); err != nil {
		return err
	}

	return nil
}
