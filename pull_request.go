package main

import (
	"fmt"

	"github.com/google/go-github/github"
)

func listPullRequest(repository *repository, client github.Client) error {
	pullRequests, _, err := client.PullRequests.List(repository.user, repository.repositoryName, nil)
	if err != nil {
		return err
	}
	for _, pullRequest := range pullRequests {
		fmt.Println(*pullRequest.Number, *pullRequest.Title)
	}
	return nil
}
