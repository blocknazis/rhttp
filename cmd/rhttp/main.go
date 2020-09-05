package main

import (
	"github.com/blocknazis/rhttp/internal/env"
	"github.com/blocknazis/rhttp/internal/updater"
	"github.com/blocknazis/rhttp/internal/webserver"
	"strings"
	"time"
)

func main() {
	// Load the optional .env file
	env.Load()

	// Start the repository updater
	repoUpdater := updater.New(env.Get("RHTTP_REPO_URL", ""), env.Duration("RHTTP_UPDATE_PERIOD", time.Hour), strings.Split(env.Get("RHTTP_BLACKLIST", ".git/"), ";;"))
	go func() {
		err := repoUpdater.ScheduleUpdates()
		if err != nil {
			panic(err)
		}
	}()

	// Start the web server
	panic(webserver.Run(env.Get("RHTTP_WEB_ADDRESS", ":8080"), repoUpdater))
}
