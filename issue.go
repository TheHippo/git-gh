package main

import (
	"fmt"

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
	fmt.Println(client)
	fmt.Println(rootFolder)
	return nil
}
