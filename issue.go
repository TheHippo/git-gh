package main

import (
	"fmt"
	"os"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	//"github.com/google/go-github/github"
)

var issueCommand = &commander.Command{
	UsageLine: "issue",
	Short:     "list & search issues",
	Long: `
List and search issues
`,
	Run:  runIssues,
	Flag: *flag.NewFlagSet("issue", flag.ExitOnError),
}

func init() {
	setRootFolderFlag(issueCommand)
}

func runIssues(cmd *commander.Command, args []string) error {
	fmt.Println("Issues...")

	client := getClient()
	repo := newRepository(rootFolder)
	fmt.Println(client)
	fmt.Println(repo)
	issues, response, err := client.Issues.ListByRepo(repo.user, repo.repositoryName, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(issues)
	fmt.Println(response)
	return nil
}
