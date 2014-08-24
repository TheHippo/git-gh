package main

import (
	"fmt"
	"os"

	"github.com/gonuts/commander"
)

var ghCommander = &commander.Command{
	UsageLine: os.Args[0] + " - commandline interface for GitHub",
}

func init() {

	ghCommander.Subcommands = []*commander.Command{
		issueCommand,
		pullRequestCommand,
	}
}

func main() {
	err := ghCommander.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	return
}
