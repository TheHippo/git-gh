package main

import (
	"github.com/gonuts/commander"
)

var rootFolder string

func setRootFolderFlag(command *commander.Command) {
	command.Flag.StringVar(&rootFolder, "git-folder", ".", "Set folder for git repository. Usually current folder")
}
