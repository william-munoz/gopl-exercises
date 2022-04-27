// Package github provides Go's API for GitHub.
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetIssues gets a certain number of issues from the specified GitHub repository.
func GetIssues(owner, repo string) ([]Issue, error) {
	req, err := http.NewRequest("GET", getIssuesURL(owner, repo), nil)
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

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}
