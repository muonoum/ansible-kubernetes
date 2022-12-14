package main

import (
	"io/fs"
	"net/http"
)

func staticHandler() http.Handler {
	fs, _ := fs.Sub(static, "embed")

	return http.FileServer(
		&FallbackFileSystem{
			FileSystem: http.FS(fs),
			Fallback:   "index.html",
		})
}
