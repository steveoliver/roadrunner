package gzip

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
)

const PluginName = "gzip"

type Plugin struct{}

// needed for the Endure
func (g *Plugin) Init() error {
	return nil
}

func (g *Plugin) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gziphandler.GzipHandler(next).ServeHTTP(w, r)
	})
}

func (g *Plugin) Name() string {
	return PluginName
}
