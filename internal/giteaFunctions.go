package internal

import (
	"fmt"
	"strings"

	"code.gitea.io/sdk/gitea"
	"github.com/google/go-github/v64/github"
)

func CreateOrgFromUser (giteaClient *gitea.Client, user *github.User) ( created_user *gitea.Organization) {
	if user.Bio == nil {
		user.Bio = new(string)
		*user.Bio = ""
	}

	if user.Blog == nil {
		user.Blog = new(string)
		*user.Blog = ""
	}

	if user.Location == nil {
		user.Location = new(string)
		*user.Location = ""
	}
	
	created_user, _, err := giteaClient.AdminCreateOrg(GiteaLogin.Username,gitea.CreateOrgOption{
		Name: *user.Login,
		FullName: *user.Name,
		Description: *user.Bio,
		Website: *user.Blog,
		Location: *user.Location,
		Visibility: "public",
	})
	CheckIfError(err)

	err = ProcessAvatar(*user.Login, *user.AvatarURL)
	CheckIfError(err)

	return created_user
}

func CreateOrgFromOrg (giteaClient *gitea.Client, org *github.Organization) ( created_org *gitea.Organization) {

	// This is hacky, but I suck at golang so this will have to do
	if org.Description == nil {
		org.Description = new(string)
		*org.Description = ""
	}

	if org.Blog == nil {
		org.Blog = new(string)
		*org.Blog = ""
	}

	if org.Location == nil {
		org.Location = new(string)
		*org.Location = ""
	}
	
	created_org, _, err := giteaClient.AdminCreateOrg(GiteaLogin.Username,gitea.CreateOrgOption{
		Name: *org.Login,
		FullName: *org.Login,
		Description: *org.Description,
		Website: *org.Blog,
		Location: *org.Location,
		Visibility: "public",
	})
	CheckIfError(err)

	err = ProcessAvatar(*org.Login, *org.AvatarURL)

	if err != nil && strings.Contains(err.Error(), "user already exists") {
		err = nil
	} else {
		CheckIfError(err)
	}

	return created_org
}

func ImportGiteaRepo (giteaClient *gitea.Client, RepoList []*github.Repository, GiteaOrg string) (err error) {
	for _, repo := range RepoList {
		fmt.Println("Importing " + *repo.Name)
		if repo.Description == nil {
			repo.Description = new(string)
			*repo.Description = ""
		}
		_, _, err := giteaClient.MigrateRepo(gitea.MigrateRepoOption{
			CloneAddr: *repo.CloneURL,
			RepoName: *repo.Name,
			RepoOwner: GiteaOrg,
			Mirror: true,
			Description: *repo.Description,
		})

		if err != nil && strings.Contains(err.Error(), "The repository with the same name already exists."){
			err = nil
		}

		CheckIfError(err)
		fmt.Println("Successfully imported " + *repo.Name)

	}
	return err

}