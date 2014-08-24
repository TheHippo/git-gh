package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

var pullRequestCommand = &commander.Command{
	UsageLine: "pullrequest",
	Short:     "list & search pull requests",
	Long: `
List and search pullrequests
`,
	Run:  runPullRequest,
	Flag: *flag.NewFlagSet("pullrequest", flag.ExitOnError),
}

func init() {
	setRootFolderFlag(pullRequestCommand)
}

func runPullRequest(cmd *commander.Command, args []string) error {
	fmt.Println("pullrequests...")
	fmt.Println(rootFolder)
	return nil
}
