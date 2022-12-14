package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"

	"haproxy-dashboard/internal/web"
)

type StartCommand struct {
	Address       string        `env:"READER_ADDRESS" required`
	Interval      time.Duration `env:"READER_INTERVAL" default:"1s"`
	ErrorInterval time.Duration `env:"READER_ERROR_INTERVAL" default:"5s"`
}

func (cmd StartCommand) Run(config Config) error {
	var cancel func()
	defer cancel()
	ctx := context.Background()
	ctx, cancel = context.WithCancel(ctx)

	client, err := web.Client(config.CACert, config.Timeout)
	if err != nil {
		return err
	}

	reader := new(Reader)
	go reader.Run(ctx, client, config, cmd)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := reader.Get()
		json.NewEncoder(w).Encode(data)
	})

	log.Info().
		Str("address", cmd.Address).
		Msg("start listener")

	return web.Server(cmd.Address, gziphandler.GzipHandler(mux)).
		ListenAndServe()
}
