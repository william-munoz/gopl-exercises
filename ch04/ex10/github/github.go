// Package github provides Go's API for GitHub.
package github

import "time"

// IssuesURL is the URL for searching for GitHub Issues.
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult represents the search results for GitHub Issue.
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue represents a GitHub Issue.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User stands for GitHub User.
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
