// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Load the GitHub API token from the environment variable
	token, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	// Initialize the GitHub client with the loaded token
	client := NewGitHubClient(token)

	// Get the username to query
	username := "your_username" // Replace with actual username or get it from user input

	// Fetch the repositories for the specified username
	repos, err := ListRepositories(client, username)
	if err != nil {
		fmt.Println("Error fetching repositories:", err)
		os.Exit(1)
	}

	// Print out the repository names
	for _, repo := range repos {
		fmt.Println("Repository Name:", repo.Name)
	}
}
