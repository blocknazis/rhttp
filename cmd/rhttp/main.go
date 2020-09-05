package main

import (
	"github.com/blocknazis/rhttp/internal/env"
	"github.com/blocknazis/rhttp/internal/updater"
	"github.com/blocknazis/rhttp/internal/webserver"
	"log"
	"strings"
	"time"
)

func main() {
	// Load the optional .env file
	env.Load()

	// Start the repository updater
	log.Println("Starting and scheduling the repository updater...")
	repoURL := env.Get("RHTTP_REPO_URL", "")
	updatePeriod := env.Duration("RHTTP_UPDATE_PERIOD", time.Hour)
	blacklist := strings.Split(env.Get("RHTTP_BLACKLIST", ".git/"), ";;")
	repoUpdater := updater.New(repoURL, updatePeriod, blacklist)
	go func() {
		err := repoUpdater.ScheduleUpdates()
		if err != nil {
			panic(err)
		}
	}()

	// Start the web server
	log.Println("Starting the web server...")
	panic(webserver.Run(env.Get("RHTTP_WEB_ADDRESS", ":8080"), repoUpdater))
}
