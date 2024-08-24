// client.go
package main

import (
	"net/http"
	"time"
)

// GitHubClient wraps the HTTP client with authentication headers for GitHub
type GitHubClient struct {
	httpClient *http.Client
	token      string
}

// NewGitHubClient initializes and returns a new GitHubClient
func NewGitHubClient(token string) *GitHubClient {
	return &GitHubClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		token: token,
	}
}

// DoRequest performs an HTTP request with the appropriate headers
func (client *GitHubClient) DoRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "token "+client.token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	return client.httpClient.Do(req)
}
