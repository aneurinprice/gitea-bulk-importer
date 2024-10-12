package internal

import (
	"context"
	"github.com/google/go-github/v66/github"
	log "github.com/sirupsen/logrus"
)

var GithubListOptions = github.ListOptions{
	PerPage: 100,
}

func GetUser(ghclient *github.Client, objectName string) (user *github.User) {
	user,_,err := ghclient.Users.Get(context.Background(), objectName)
	CheckIfError(err)
	return user
}

func GetOrg(ghclient *github.Client, objectName string) (org *github.Organization) {
	org,_,err := ghclient.Organizations.Get(context.Background(), objectName)
	CheckIfError(err)
	return org
}

func GetGithubUserRepos(ghclient *github.Client, user *github.User) (repoList []*github.Repository) {
	UserListOptions := github.RepositoryListByUserOptions{ListOptions: GithubListOptions}
	for {
		log.Debug("Parsing page: ", UserListOptions.Page)
		repoPage, resp, err := ghclient.Repositories.ListByUser(context.Background(), Args.Input, &UserListOptions)
		CheckIfError(err)
		repoList = append(repoList, repoPage...)
		log.Debug("RepoList: ", repoList)
		if resp.LastPage == 0 {break}
		UserListOptions.ListOptions.Page++
	}
	return repoList
}

func GetGithubOrgRepos(ghclient *github.Client, org *github.Organization) (repoList []*github.Repository) {
	OrgListOptions := github.RepositoryListByOrgOptions{ListOptions: GithubListOptions}
	log.Debug("Getting repos for org: ", org.Login)
	log.Debug("Using Settings: ", OrgListOptions)
	for {
		log.Debug("Parsing page: ", OrgListOptions.Page)
		repoPage, resp, err := ghclient.Repositories.ListByOrg(context.Background(), Args.Input, &OrgListOptions)
		CheckIfError(err)
		repoList = append(repoList, repoPage...)
		log.Debug("RepoList: ", repoList)
		if resp.LastPage == 0 {break}
		OrgListOptions.ListOptions.Page++
	}
	return repoList
}

func FilterRepoList(repoList []*github.Repository) (filteredRepoList []*github.Repository) {
	droppedRepos := 0
	for _, repo := range repoList {
		if !Args.IncludeForks && *repo.Fork {
			log.Warn("Skipping forked repo: ", *repo.FullName)
			droppedRepos++
		} else {
			filteredRepoList = append(filteredRepoList, repo)
			log.Debug("Adding repo: ", *repo.FullName)
		}
	}
	log.Debug("Before filtering: ", len(repoList))
	log.Debug("Dropped repos: ", droppedRepos)
	log.Debug("After filtering: ", len(filteredRepoList))
	log.Info("Unaccounted for: ", len(repoList) - len(filteredRepoList) - droppedRepos)
	
	return filteredRepoList
}