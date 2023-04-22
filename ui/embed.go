// Package ui handles the PocketBase Admin frontend embedding.
package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:dist
var distDir embed.FS

// GetFileSystem returns the embedded frontend files.
func GetFileSystem() http.FileSystem {
	stripped, err := fs.Sub(distDir, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(stripped)
}