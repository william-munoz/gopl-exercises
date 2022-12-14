// Package github provides Go's API for GitHub.
package github

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getIssueURL(owner, repo, number string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, number)
}

func getIssuesURL(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
}

func setAuthorization(req *http.Request) error {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return fmt.Errorf("GITHUB_TOKEN is not set")
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	return nil
}

// Issue represents a GitHub Issue.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	Comments  int
	UpdatedAt time.Time `json:"updated_at"`
}

// User represents a user on GitHub.
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
