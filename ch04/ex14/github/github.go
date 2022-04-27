// Package github provides Go's API for GitHub.
package github

import (
	"fmt"
	"time"
)

func getIssuesURL(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?state=all", owner, repo)
}

// Issue represents a GitHub Issue.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	Assignees []*User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	UpdatedAt time.Time `json:"updated_at"`
	Milestone *Milestone
}

// User represents a user on GitHub.
type User struct {
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
	ID        int
	Login     string
}

// Milestone represents a milestone on GitHub.
type Milestone struct {
	Description string
	HTMLURL     string `json:"html_url"`
	ID          int
	State       string
	Title       string
}

// Equals returns whether it can be considered equal to a given User.
func (u *User) Equals(x *User) bool {
	return u.ID == x.ID
}

// Equals returns whether it can be considered equal to a given Milestone.
func (m *Milestone) Equals(x *Milestone) bool {
	return m.ID == x.ID
}
