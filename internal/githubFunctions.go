package internal

import (
	"context"
	"github.com/google/go-github/v64/github"
)

func GetUser(github *github.Client, objectName string) (user *github.User) {
	user,_,err := github.Users.Get(context.Background(), objectName)
	CheckIfError(err)
	return user
}

func GetOrg(github *github.Client, objectName string) (org *github.Organization) {
	org,_,err := github.Organizations.Get(context.Background(), objectName)
	CheckIfError(err)
	return org
}

func GetGithubUserRepos(github *github.Client, user *github.User) (repoList []*github.Repository) {
	repoList, _, err := github.Repositories.ListByUser(context.Background(), *user.Login, nil)
	CheckIfError(err)
	return repoList
}

func GetGithubOrgRepos(github *github.Client, org *github.Organization) (repoList []*github.Repository) {
	repoList, _, err := github.Repositories.ListByOrg(context.Background(), *org.Login, nil)
	CheckIfError(err)
	return repoList
}