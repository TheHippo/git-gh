package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const (
	appVersion = "0.0.1"

	rootDirectoryFlag = "git-directory"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-gh"
	app.Usage = "github command line tools"
	app.Version = appVersion

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, gd", rootDirectoryFlag),
			Value: ".",
			Usage: "Root of git repository",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "issue",
			ShortName: "i",
			Usage:     "List and create issues",
			Subcommands: []cli.Command{
				{
					Name:      "list",
					ShortName: "l",
					Usage:     "List issues",
					Action: func(c *cli.Context) {
						fmt.Println("List issues")
						fmt.Println(c.GlobalString(rootDirectoryFlag))
					},
				},
				{
					Name:      "create",
					ShortName: "c",
					Usage:     "Create issue",
					Action: func(c *cli.Context) {
						fmt.Println("Create issue")
					},
				},
			},
		},
		{
			Name:      "pullrequest",
			ShortName: "p",
			Usage:     "List pull request",
			Action: func(c *cli.Context) {
				fmt.Println("Pull request")
			},
		},
	}

	app.Run(os.Args)
}
