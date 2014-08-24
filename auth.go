package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
)

const (
	tokenEnv = "GIT_GH_TOKEN"
)

var authToken string

func init() {
	authToken := os.Getenv(tokenEnv)
	if len(authToken) == 0 {
		fmt.Println("Could find auth token in enviroment")
		os.Exit(1)
	}
}

func getClient() github.Client {
	return *github.NewClient()
}
