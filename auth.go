package main

import (
	"fmt"
	"os"

	"code.google.com/p/goauth2/oauth"
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

func getTransport() *oauth.Transport {
	return &oauth.Transport{
		Token: &oauth.Token{
			AccessToken: authToken,
		},
	}
}

func getClient() github.Client {
	return *github.NewClient(getTransport().Client())
}
