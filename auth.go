package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

const (
	tokenEnv       = "GIT_GH_TOKEN"
	githubPassword = "x-oauth-basic"
)

var authToken string

type transport struct {
	token     string
	Transport http.RoundTripper
}

func cloneRequest(req *http.Request) *http.Request {
	// shallow copy of the struct
	clone := new(http.Request)
	*clone = *req
	// deep copy of the Header
	clone.Header = make(http.Header)
	for k, s := range req.Header {
		clone.Header[k] = s
	}
	return clone
}

func (t *transport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req)
	req.SetBasicAuth(t.token, githubPassword)
	return t.transport().RoundTrip(req)
}

func createHTTPClient(token string) *http.Client {
	return &http.Client{
		Transport: &transport{
			token: token,
		},
	}
}

func init() {
	// TODO: fetch this from a safer location
	authToken = os.Getenv(tokenEnv)
	if len(authToken) == 0 {
		fmt.Println("Could find auth token in enviroment")
		os.Exit(1)
	}
}

func getClient() github.Client {
	return *github.NewClient(createHTTPClient(authToken))
}
