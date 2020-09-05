package env

import (
	"github.com/joho/godotenv"
	"os"
	"time"
)

// Load loads an optional .env file
func Load() {
	godotenv.Load()
}

// Get returns the content of the environment variable with the given key or the given fallback
func Get(key, fallback string) string {
	found := os.Getenv(key)
	if found == "" {
		return fallback
	}
	return found
}

// Duration uses Get and parses it into a duration
func Duration(key string, fallback time.Duration) time.Duration {
	parsed, _ := time.ParseDuration(Get(key, fallback.String()))
	return parsed
}
