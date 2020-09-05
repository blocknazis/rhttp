package webserver

import (
	"github.com/blocknazis/rhttp/internal/updater"
	"github.com/valyala/fasthttp"
)

// Run runs the web server on the specified address
func Run(address string, repoUpdater *updater.Updater) error {
	return fasthttp.ListenAndServe(address, handler(repoUpdater))
}
