package api

import (
	"io/fs"
	"net/http"
	"path"
	"strings"

	"marketplace/apps/web"
)

// spa serves the built React app from the binary.
//
// Embedding it means production is one container on one port, which is how this
// gets deployed. In development Vite serves the app instead and proxies /api
// here, so this handler simply reports that nothing was built.
func (s *Server) spa() http.Handler {
	dist, err := web.Dist()
	if err != nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api/") {
				writeErr(w, http.StatusNotFound, "not_found", "No such endpoint.")
				return
			}
			writeJSON(w, http.StatusOK, map[string]string{
				"message": "The web app is not embedded in this build. Run the Vite dev server, or build it first.",
			})
		})
	}

	files := http.FileServer(http.FS(dist))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			writeErr(w, http.StatusNotFound, "not_found", "No such endpoint.")
			return
		}

		// A path with no extension is a client-side route, so hand back the app
		// shell and let the router work it out.
		clean := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		if clean == "" || path.Ext(clean) == "" {
			serveIndex(w, r, dist)
			return
		}
		if _, err := fs.Stat(dist, clean); err != nil {
			serveIndex(w, r, dist)
			return
		}
		files.ServeHTTP(w, r)
	})
}

func serveIndex(w http.ResponseWriter, r *http.Request, dist fs.FS) {
	body, err := fs.ReadFile(dist, "index.html")
	if err != nil {
		writeErr(w, http.StatusNotFound, "not_found", "The web app is not available.")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	_, _ = w.Write(body)
}
