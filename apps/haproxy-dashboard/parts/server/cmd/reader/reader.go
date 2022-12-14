package main

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type Reader struct {
	mutex sync.RWMutex
	rows  []map[string]string
}

func (r *Reader) Get() []map[string]string {
	defer r.mutex.RUnlock()
	r.mutex.RLock()
	return r.rows
}

func (r *Reader) Put(rows []map[string]string) {
	defer r.mutex.Unlock()
	r.mutex.Lock()
	r.rows = rows
}

func (r *Reader) Run(ctx context.Context, client *http.Client, config Config, cmd StartCommand) {
	log.Info().
		Str("source", config.SourceURL).
		Dur("interval", cmd.Interval).
		Dur("error-interval", cmd.ErrorInterval).
		Msg("start reader")

	timer := time.NewTimer(0)

	for {
		select {
		case <-ctx.Done():
			return

		case <-timer.C:
			rows, err := readURL(client, config.SourceURL)
			if err != nil {
				log.Error().Err(err).
					Str("url", config.SourceURL).
					Msg("could not read from url")
				timer.Reset(cmd.Interval)
				continue
			}

			r.Put(removeEmpty(rows))
			timer.Reset(cmd.Interval)
		}
	}
}

func removeEmpty(rows []map[string]string) []map[string]string {
	for index, row := range rows {
		for key, value := range row {
			if value == "" {
				delete(row, key)
			}

			rows[index] = row
		}
	}

	return rows
}

func readURL(client *http.Client, url string) ([]map[string]string, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return read(response.Body)
}

func read(reader io.ReadCloser) ([]map[string]string, error) {
	defer reader.Close()

	rows := make([]map[string]string, 0)
	skip := make([]byte, 2)
	if n, err := reader.Read(skip); err != nil {
		return rows, err
	} else if n != len(skip) {
		return rows, errors.New("could not skip header comment")
	}

	records := csv.NewReader(reader)
	var headers []string
	headers, err := records.Read()
	if err != nil {
		return rows, nil
	}

	for {
		record, err := records.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return rows, err
		}

		values := make(map[string]string)
		for i := range headers {
			if headers[i] != "" {
				values[headers[i]] = record[i]
			}
		}

		rows = append(rows, values)
	}

	return rows, nil
}
