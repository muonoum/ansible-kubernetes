package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func maybeNotify(method string, url string) {
	if url != "" {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			log.Warn().Err(err).
				Str("url", url).
				Msg("notify request failed")
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Warn().Err(err).
				Str("url", url).
				Msg("notify request failed")
			return
		}

		if res.StatusCode >= 400 {
			log.Warn().
				Int("status-code", res.StatusCode).
				Msg("notify request failed")
			return
		}

		log.Info().
			Str("method", method).
			Str("url", url).
			Int("response", res.StatusCode).
			Msg("notify")
	}
}
