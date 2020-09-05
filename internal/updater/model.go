package updater

import "time"

// Updater represents a repository updater
type Updater struct {
	RepoURL    string
	Period     time.Duration
	LastUpdate time.Time
	State      State
}

// State represents the state of a repository updater
type State int

// These constants represent the different states of a repository updater
const (
	StateAvailable = 0
	StateUpdating  = 1
)
