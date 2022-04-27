// ch04 / ex14 is a web server that allows you to browse bug reports, milestones, and a list of users with a single contact to GitHub.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/williammunozr/gopl-exercises/ch04/ex14/github"
)

var navigation = `
<p>
<a href='/'>Issues</a> /
<a href='/milestones'>Milestones</a> /
<a href='/users'>Users</a>
</p>
`
var issuesTemplate = template.Must(template.New("issues").Parse(navigation + `
<h1>{{len .}} issue{{if ne (len .) 1}}s{{end}}</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var milestonesTemplate = template.Must(template.New("milestones").Parse(navigation + `
<h1>{{len .}} milestone{{if ne (len .) 1}}s{{end}}</h1>
<table>
<tr style='text-align: left'>
<th>Title</th>
<th>State</th>
</tr>
{{range .}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	<td>{{.State}}</td>
</tr>
{{end}}
</table>
`))

var usersTemplate = template.Must(template.New("users").Parse(navigation + `
<h1>{{len .}} user{{if ne (len .) 1}}s{{end}}</h1>
<table>
<tr style='text-align: left'>
<th>Avatar</th>
<th>Username</th>
</tr>
{{range .}}
<tr>
	<td><a href='{{.HTMLURL}}'><img src='{{.AvatarURL}}' width='32' height='32'></td>
	<td><a href='{{.HTMLURL}}'>{{.Login}}</td>
</tr>
{{end}}
</table>
`))

var issues []github.Issue
var milestones []github.Milestone
var users []github.User

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: ex14 OWNER REPO")
		os.Exit(1)
	}

	owner, repo := os.Args[1], os.Args[2]
	err := generateCache(owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handleIssues)
	http.HandleFunc("/milestones", handleMilestones)
	http.HandleFunc("/users", handleUsers)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generateCache(owner, repo string) error {
	got, err := github.GetIssues(owner, repo)
	if err != nil {
		return err
	}

	issues = got

	for _, issue := range issues {
		if issue.Milestone != nil {
			milestones = appendMilestoneAsSet(milestones, issue.Milestone)
		}
		for _, assignee := range issue.Assignees {
			users = appendUserAsSet(users, assignee)
		}
		users = appendUserAsSet(users, issue.User)
	}
	return nil
}

func handleIssues(w http.ResponseWriter, r *http.Request) {
	issuesTemplate.Execute(w, issues)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	usersTemplate.Execute(w, users)
}

func handleMilestones(w http.ResponseWriter, r *http.Request) {
	milestonesTemplate.Execute(w, milestones)
}

// appendMilestoneAsSet adds the given milestone milestone to the given milestone array set.
// However, if set already contains milestones, do not add milestones.
func appendMilestoneAsSet(set []github.Milestone, milestone *github.Milestone) []github.Milestone {
	if !includesMilestone(set, milestone) {
		return append(set, *milestone)
	}
	return set
}

// includesMilestone returns whether the array of given milestones contains the given milestone milestones.
func includesMilestone(array []github.Milestone, milestone *github.Milestone) bool {
	for _, value := range array {
		if value.Equals(milestone) {
			return true
		}
	}
	return false
}

// appendUserAsSet adds the given user user to the given user array set.
// However, if set already contains user, do not add user.
func appendUserAsSet(set []github.User, user *github.User) []github.User {
	if !includesUser(set, user) {
		return append(set, *user)
	}
	return set
}

// includesUser returns whether the array array of given users contains the given user user.
func includesUser(array []github.User, user *github.User) bool {
	for _, value := range array {
		if value.Equals(user) {
			return true
		}
	}
	return false
}
