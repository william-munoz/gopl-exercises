// Package github provides Go's API for GitHub.
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// patchIssue performs write operations on GitHub Issue.
func patchIssue(owner, repo, number string, fields map[string]string) error {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", getIssueURL(owner, repo, number), buf)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Content-Type", "application/json")
	err = setAuthorization(req)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// All paths below this line should close resp.Body.
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("patch issue failed: %s", resp.Status)
	}

	return nil
}

// UpdateIssue updates the GitHub Issue.
func UpdateIssue(owner, repo, number string, fields map[string]string) error {
	return patchIssue(owner, repo, number, fields)
}

// ReopenIssue reopens the closed GitHub Issue.
func ReopenIssue(owner, repo, number string) error {
	fields := map[string]string{
		"state": "open",
	}
	return patchIssue(owner, repo, number, fields)
}

// CloseIssue edits a GitHub Issue.
func CloseIssue(owner, repo, number string) error {
	fields := map[string]string{
		"state": "closed",
	}
	return patchIssue(owner, repo, number, fields)
}
