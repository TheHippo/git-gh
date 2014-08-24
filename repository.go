package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type repository struct {
	folder         string
	user           string
	repositoryName string
}

// https://github.com/gonuts/commander.git
// git@github.com:gonuts/commander.git

var matchGithubURL = regexp.MustCompile(`(?i)github.com[:/]([a-z.0-9]+)/([a-z.0-9]+)\.git`)

func findGithubRemote(path string) (string, string, error) {
	// git --git-dir=$PATH/.git --work-tree=$PATH/ remote -v
	remoteCommand := exec.Command("git", fmt.Sprintf("--git-dir=%s/.git", path), fmt.Sprintf("--work-tree=%s/", path), "remote", "-v")
	remoteCommandOutput, err := remoteCommand.Output()
	if err != nil {
		return "", "", err
	}
	lines := strings.Split(string(remoteCommandOutput), "\n")
	for _, line := range lines {
		matches := matchGithubURL.FindAllStringSubmatch(line, -1)
		if len(matches) > 0 && len(matches[0]) == 3 {
			return matches[0][1], matches[0][2], nil
		}
	}
	return "", "", fmt.Errorf("Could not find user name and repository name")
}

func findRepositoryRoot(path string) (string, error) {

	if !strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s/", path)
	}

	// This actually does not find the root of the repository, but throws an error if folder is not
	// repository root
	// git --git-dir=$PATH/.git --work-tree=$PATH/ rev-parse --show-toplevel
	rootPathCommand := exec.Command("git", fmt.Sprintf("--git-dir=%s.git", path), fmt.Sprintf("--work-tree=%s", path), "rev-parse", "--show-toplevel")
	rootPathCommandOutput, err := rootPathCommand.Output()
	if err != nil {
		return "", err
	}
	fmt.Println(strings.TrimSpace(string(rootPathCommandOutput)))
	return strings.TrimSpace(string(rootPathCommandOutput)), nil
}

func newRepository(path string) *repository {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error determing root folder path: ", err.Error())
		os.Exit(1)
		return &repository{}
	}
	rootPath, err := findRepositoryRoot(absPath)
	if err != nil {
		fmt.Println("Could not find git repository root folder: ", err.Error())
		os.Exit(1)
		return &repository{}
	}
	user, repositoryName, err := findGithubRemote(rootPath)
	if err != nil {
		fmt.Println("Could not determine github repository: ", err.Error())
		os.Exit(1)
		return &repository{}
	}
	return &repository{
		folder:         rootPath,
		user:           user,
		repositoryName: repositoryName,
	}
}
