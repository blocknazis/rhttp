package updater

import (
	"os"
	"os/exec"
	"time"
)

// New creates a new repository updater
func New(repoURL string, period time.Duration) *Updater {
	return &Updater{
		RepoURL: repoURL,
		Period:  period,
		State:   StateUpdating,
	}
}

// ScheduleUpdates schedules the repository updates
func (updater *Updater) ScheduleUpdates() error {
	// Update the git repository
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

	// Set the state of the updater to 'available'
	updater.State = StateAvailable
	return nil
}
