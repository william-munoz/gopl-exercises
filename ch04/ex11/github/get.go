// Package github provides Go's API for GitHub.
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetIssue gets the specified GitHub Issue.
func GetIssue(owner, repo, number string) (*Issue, error) {
	req, err := http.NewRequest("GET", getIssueURL(owner, repo, number), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// All paths below this line should close resp.Body.
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
