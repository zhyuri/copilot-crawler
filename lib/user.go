package lib

import (
	"context"
	"github.com/shurcooL/githubv4"
)

type GithubUser struct {
	ID         githubv4.String
	Name       githubv4.String
	Login      githubv4.String
	Bio        githubv4.String
	Email      githubv4.String
	AvatarURL  githubv4.URI
	URL        githubv4.URI
	WebsiteUrl githubv4.URI
}

func NewGithubUser(name string) (GithubUser, error) {
	var query struct {
		GithubUser `graphql:"user(login:$name)"`
	}
	param := map[string]interface{}{
		"name": githubv4.String(name),
	}

	err := client().Query(context.Background(), &query, param)
	log.Debug("query github user with param %v got %+v, error %v", param, query, err)
	if err != nil {
		log.Fatal("query github user got error %v, param %v, query %+v", param, query, err)
		return GithubUser{}, err
	}
	return query.GithubUser, nil
}
