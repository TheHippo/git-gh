package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
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

func runIssues(cmd *commander.Command, args []string) error {
	fmt.Println("Issues...")
	return nil
}
