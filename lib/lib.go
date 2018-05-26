package lib

import (
	"context"
	"github.com/op/go-logging"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
)

var Log *logging.Logger
var log *logging.Logger
var logFormat = logging.MustStringFormatter(
	`%{color} %{level:.4s} %{shortfile} %{callpath} ▶ %{color:reset} %{message}`,
)

func init() {

	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, logFormat)
	backendFormatted := logging.AddModuleLevel(backendFormatter)
	logging.SetBackend(backend, backendFormatted)

	log = logging.MustGetLogger("copilot")
	Log = log
}

var githubClient *githubv4.Client

func client() *githubv4.Client {
	if githubClient != nil {
		return githubClient
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	githubClient = githubv4.NewClient(httpClient)
	return githubClient
}
