// config.go
package main

import (
	"errors"
	"os"
)

// LoadConfig loads the GitHub API token from environment variables
func LoadConfig() (string, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return "", errors.New("GitHub token not set in environment variables")
	}
	return token, nil
}
