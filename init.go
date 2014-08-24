package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

var initCommand = &commander.Command{
	UsageLine: "init",
	Short:     "set up git gh",
	Long: `
Set up git gh
`,
	Run:  runInit,
	Flag: *flag.NewFlagSet("init", flag.ExitOnError),
}

func runInit(cmd *commander.Command, args []string) error {
	fmt.Println("Init...")
	return nil
}
