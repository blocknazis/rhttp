package webserver

import (
	"github.com/blocknazis/rhttp/internal/updater"
	"github.com/valyala/fasthttp"
)

// handler handles all incoming web requests
func handler(repoUpdater *updater.Updater) fasthttp.RequestHandler {
	// Define the file server
	fs := &fasthttp.FS{
		Root:               "./data",
		GenerateIndexPages: true,
		PathNotFound: func(ctx *fasthttp.RequestCtx) {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString("File not found.")
		},
	}
	fsHandler := fs.NewRequestHandler()

	// Return the handler function
	return func(ctx *fasthttp.RequestCtx) {
		// Check if the updater is in updating state
		if repoUpdater.State == updater.StateUpdating {
			ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
			ctx.SetBodyString("Repository is being updated...")
			return
		}

		// Serve the files otherwise
		fsHandler(ctx)
	}
}
