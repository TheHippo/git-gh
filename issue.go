package main

import (
	"fmt"

	"github.com/google/go-github/github"
)

func listIssues(repository *repository, client github.Client) error {
	issues, _, err := client.Issues.ListByRepo(repository.user, repository.repositoryName, nil)
	if err != nil {
		return err
	}
	for _, issue := range issues {
		fmt.Println(*issue.Number, *issue.Title)
	}
	return nil
}
