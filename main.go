package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

const (
	appVersion = "0.0.1"

	rootDirectoryFlag = "git-directory"
)

const (
	errorCodeIssueList = iota + 1
	errorCodeIssueCreate
	errorCodePullRequestList
)

func getBase(path string) (repository *repository, client github.Client) {
	repository = newRepository(path)
	client = getClient()
	return
}

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
						repository, client := getBase(c.GlobalString(rootDirectoryFlag))
						if err := listIssues(repository, client); err != nil {
							fmt.Println("Could not list issues: ", err.Error())
							os.Exit(errorCodeIssueList)
						}

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
			Usage:     "List and create pull request",
			Subcommands: []cli.Command{
				{
					Name:      "list",
					ShortName: "l",
					Usage:     "List pull requests",
					Action: func(c *cli.Context) {
						repository, client := getBase(c.GlobalString(rootDirectoryFlag))
						if err := listPullRequest(repository, client); err != nil {
							fmt.Println("Could not list pull requests: ", err.Error())
							os.Exit(errorCodePullRequestList)
						}
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
