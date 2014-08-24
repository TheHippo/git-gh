package main

import (
	"os"
	"text/template"

	"github.com/google/go-github/github"
)

var listIssuesTemplate = template.Must(template.New("listIssues").Parse(`
{{ range . }}{{ .Number }} {{ .Title }}
{{ end }}
`))

func stateString(closed bool) string {
	if closed {
		return "closed"
	}
	return "open"
}

func listIssues(repository *repository, client github.Client, closed bool) error {
	query := &github.IssueListByRepoOptions{
		State: stateString(closed),
	}
	issues, _, err := client.Issues.ListByRepo(repository.user, repository.repositoryName, query)
	if err != nil {
		return err
	}
	err = listIssuesTemplate.Execute(os.Stdout, issues)
	return err
}
