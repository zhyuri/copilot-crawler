package lib

import (
	"context"
	"github.com/shurcooL/githubv4"
)

type GithubRepo struct {
	ID               githubv4.String
	Name             githubv4.String
	Login            githubv4.String
	Bio              githubv4.String
	DefaultBranchRef githubv4.String
	AvatarURL        githubv4.URI
	URL              githubv4.URI
	WebsiteUrl       githubv4.URI
}

func NewGithubRepo(name string) (GithubRepo, error) {
	var query struct {
		GithubRepo `graphql:"repositoryinfo(login:$name)"`
	}
	param := map[string]interface{}{
		"name": githubv4.String(name),
	}

	err := client().Query(context.Background(), &query, param)
	log.Debug("query hub repo with param %v got %+v, error %v", param, query, err)
	if err != nil {
		log.Fatal("query hub repo got error %v, param %v, query %+v", param, query, err)
		return GithubRepo{}, err
	}
	return query.GithubRepo, nil
}
