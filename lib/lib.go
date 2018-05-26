package lib

import (
	"context"
	"github.com/op/go-logging"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
)

var log *logging.Logger

// Log is exported for cli use
var Log *logging.Logger

func init() {
	logging.SetFormatter(logging.MustStringFormatter(
		`%{level:.4s} %{shortfile} %{callpath} ▶ %{message}`,
	))
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
