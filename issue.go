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
	//Run:  runIssues,
	Flag: *flag.NewFlagSet("issue", flag.ExitOnError),
}

var issueListCommand = &commander.Command{
	UsageLine: "list",
	Short:     "list issues for repository",
	Long: `
List all issues in this repository
`,
	Run:  runIssueList,
	Flag: *flag.NewFlagSet("issueList", flag.ExitOnError),
}

func init() {
	setRootFolderFlag(issueCommand)
	setRootFolderFlag(issueListCommand)
	issueCommand.Subcommands = []*commander.Command{
		issueListCommand,
	}
	issueListCommand.Parent = issueCommand
}

func runIssueList(cmd *commander.Command, args []string) error {
	fmt.Println("Issues...")

	client := getClient()
	repo := newRepository(rootFolder)
	issues, _, err := client.Issues.ListByRepo(repo.user, repo.repositoryName, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(len(issues))
	for _, issue := range issues {
		fmt.Println(*issue.Number, *issue.Title)
	}
	return nil
}

func runIssues(cmd *commander.Command, args []string) error {
	return nil
}
