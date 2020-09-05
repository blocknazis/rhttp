package updater

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// New creates a new repository updater
func New(repoURL string, period time.Duration, blacklist []string) *Updater {
	return &Updater{
		RepoURL:   repoURL,
		Period:    period,
		Blacklist: blacklist,
		State:     StateUpdating,
	}
}

// ScheduleUpdates schedules the repository updates
func (updater *Updater) ScheduleUpdates() error {
	// Update the git repository
	log.Println("Updating data...")
	err := updater.Update()
	if err != nil {
		return err
	}

	// Wait for the defined period
	time.Sleep(updater.Period)

	// Recursively call this method again
	return updater.ScheduleUpdates()
}

// Update updates the repository
func (updater *Updater) Update() error {
	// Set the state of the updater to 'updating'
	updater.State = StateUpdating

	// Remove the data folder
	err := os.RemoveAll("./data")
	if err != nil {
		return err
	}

	// Clone the git repository into the data folder
	cmd := exec.Command("git", "clone", updater.RepoURL, "./data")
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Remove all the blacklisted files and folders
	for _, blacklistedPath := range updater.Blacklist {
		err = os.RemoveAll(filepath.Join("./data", blacklistedPath))
		if err != nil {
			return err
		}
	}

	// Set the state of the updater to 'available'
	updater.State = StateAvailable
	return nil
}
