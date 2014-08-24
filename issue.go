package main

import (
	"fmt"

	"github.com/google/go-github/github"
)

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
	for _, issue := range issues {
		fmt.Println(*issue.Number, *issue.Title)
	}
	return nil
}
