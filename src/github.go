// github.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Repository represents a GitHub repository
type Repository struct {
	Name string `json:"name"`
}

// ListRepositories fetches repositories for a given GitHub username
func ListRepositories(client *GitHubClient, username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repositories: %d", resp.StatusCode)
	}

	var repos []Repository
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}
