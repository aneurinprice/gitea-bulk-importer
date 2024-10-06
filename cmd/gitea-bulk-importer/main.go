package main

import (
	"strings"

	"code.gitea.io/sdk/gitea"
	"github.com/aneurinprice/gitea-bulk-importer/internal"
	"github.com/google/go-github/v64/github"
)

func main() {

	err := internal.Init()
	internal.CheckIfError(err)

	githubClient := github.NewClient(nil).WithAuthToken(internal.GithubLogin.Password)
	giteaClient, _ := gitea.NewClient(internal.GiteaLogin.GiteaUrl, gitea.SetToken(internal.GiteaLogin.Password))

	switch strings.ToLower(internal.Args.Type) {
	case "user":
		user := internal.GetUser(githubClient, internal.Args.Input)
		internal.CreateOrgFromUser(giteaClient, user)
		repoList := internal.GetGithubUserRepos(githubClient, user)
		internal.ImportGiteaRepo(giteaClient, repoList, *user.Login)
	case "org":
		org := internal.GetOrg(githubClient, internal.Args.Input)
		internal.CreateOrgFromOrg(giteaClient, org)
		repoList := internal.GetGithubOrgRepos(githubClient, org)
		internal.ImportGiteaRepo(giteaClient, repoList, *org.Login)
	}
} 