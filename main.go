package main

import (
	"fmt"
	"os"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

var gh_commander = &commander.Command{
	UsageLine: os.Args[0] + " - commandline interface for GitHub",
}

var cmd1 = &commander.Command{
	Run:       exec_cmd1,
	UsageLine: "cmd1 does one thing",
	Long: `
Runs cmd 1 and exits

foobar 
`,
	Flag: *flag.NewFlagSet("cmd1", flag.ExitOnError),
}

func init() {
	gh_commander.Subcommands = []*commander.Command{
		cmd1,
	}
}

func exec_cmd1(cmd *commander.Command, args []string) error {
	fmt.Println("Hello form cmd1")
	return nil
}

func main() {
	err := gh_commander.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	return
}
