package lib

import (
	"context"
	"github.com/shurcooL/githubv4"
)

type GithubRepo struct {
	ID   githubv4.String
	Name githubv4.String
	URL  githubv4.URI
}

func NewGithubRepo(name string, owner string) (GithubRepo, error) {
	var query struct {
		GithubRepo `graphql:"repository(name:$name, owner:$owner)"`
	}
	param := map[string]interface{}{
		"name":  githubv4.String(name),
		"owner": githubv4.String(owner),
	}

	err := client().Query(context.Background(), &query, param)
	log.Debug("query hub repo with param %v got %+v, error %v", param, query, err)
	if err != nil {
		log.Fatal("query hub repo got error %v, param %v, query %+v", param, query, err)
		return GithubRepo{}, err
	}
	return query.GithubRepo, nil
}
