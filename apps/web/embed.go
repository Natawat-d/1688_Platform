// Package web carries the built storefront so the server ships as one binary.
//
// The production image runs a single container behind Caddy, which is how this
// project is deployed, so the compiled React app travels inside the Go binary
// rather than in a second web server. During development Vite serves the app and
// proxies the API, and Dist reports that nothing was embedded.
package web

import (
	"embed"
	"errors"
	"io/fs"
)

// The all: prefix keeps dot-files, which is what lets the placeholder inside
// dist satisfy the embed pattern on a checkout that has never run a web build.
//
//go:embed all:dist
var dist embed.FS

// ErrNotBuilt means the web app has not been built into this binary.
var ErrNotBuilt = errors.New("web: the app was not built into this binary")

// Dist returns the built app rooted at dist, or ErrNotBuilt when the directory
// holds nothing but its placeholder.
func Dist() (fs.FS, error) {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		return nil, err
	}
	if _, err := fs.Stat(sub, "index.html"); err != nil {
		return nil, ErrNotBuilt
	}
	return sub, nil
}
