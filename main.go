package main

import (
	"fmt"
	"os"

	"github.com/gonuts/commander"
)

var gh_commander = &commander.Command{
	UsageLine: os.Args[0] + " - commandline interface for GitHub",
}

func init() {
	gh_commander.Subcommands = []*commander.Command{
		initCommand,
		issueCommand,
		pullRequestCommand,
	}
}

func main() {
	err := gh_commander.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	return
}
