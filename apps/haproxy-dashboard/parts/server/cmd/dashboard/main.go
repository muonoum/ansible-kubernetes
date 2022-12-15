package main

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/alecthomas/kong"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"

	"haproxy-dashboard/internal/service"
	"haproxy-dashboard/internal/web"
)

func main() {
	var config Config
	cli := kong.Parse(&config)
	cli.FatalIfErrorf(cli.Validate())

	service.ConfigureLogger(config.LogFormat, config.LogLevel)
	ctx, cancel := service.Start(context.Background())

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", staticHandler())

	if config.Source != "" {
		log.Info().
			Str("source", config.Source).
			Dur("timeout", config.Timeout).
			Msg("start proxy")

		url, err := url.Parse(config.Source)
		cli.FatalIfErrorf(err)
		proxy := web.Proxy(nil, url, config.Timeout, time.Hour*24)
		mux.Handle("/stats/", proxy)
	}

	server := web.Server(config.Address, gziphandler.GzipHandler(mux))
	go func() {
		log.Info().
			Str("address", config.Address).
			Msg("start listener")

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("server error")
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("server error")
	}
}
