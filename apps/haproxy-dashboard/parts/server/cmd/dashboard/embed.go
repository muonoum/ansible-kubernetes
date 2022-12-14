package main

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
)

//go:embed embed
var static embed.FS

type FallbackFileSystem struct {
	http.FileSystem
	Fallback string
}

func (ffs *FallbackFileSystem) Open(name string) (http.File, error) {
	f, err := ffs.FileSystem.Open(name)
	if errors.Is(err, fs.ErrNotExist) {
		return ffs.FileSystem.Open(ffs.Fallback)
	}

	return f, err
}
