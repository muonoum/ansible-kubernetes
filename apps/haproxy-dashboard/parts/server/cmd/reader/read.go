package main

import (
	"github.com/k0kubun/pp/v3"
	"github.com/rs/zerolog/log"

	"haproxy-dashboard/internal/web"
)

type ReadCommand struct{}

func (cmd ReadCommand) Run(config Config) error {
	client, err := web.Client(config.CACert, config.Timeout)
	if err != nil {
		return err
	}

	rows, err := readURL(client, config.SourceURL)
	if err != nil {
		return err
	}

	log.Info().Int("rows", len(rows)).Msg("records")

	pp.Println(rows)

	return nil
}
