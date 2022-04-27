// ch04 / ex10 displays a table of GitHub Issue that matches the search term, categorized by the period when the issue was created.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kdama/gopl/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	beforeMonth := now.AddDate(0, -1, 0)
	beforeYear := now.AddDate(-1, 0, 0)

	// Report issues created in less than a month.
	fmt.Println("\n-- created at less than a month --")
	for _, item := range result.Items {
		if item.CreatedAt.After(beforeMonth) {
			printIssue(item)
		}
	}

	// Report issues created in more than a month and less than a year.
	fmt.Println("\n-- created at less than a year --")
	for _, item := range result.Items {
		if (item.CreatedAt.Before(beforeMonth) ||
			item.CreatedAt.Equal(beforeMonth)) &&
			item.CreatedAt.After(beforeYear) {
			printIssue(item)
		}
	}

	// Report issues that have been created for over a year.
	fmt.Println("\n-- created at more than a year --")
	for _, item := range result.Items {
		if item.CreatedAt.Before(beforeYear) ||
			item.CreatedAt.Equal(beforeYear) {
			printIssue(item)
		}
	}
}

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
}
