package internal

import (
	"strings"
	"code.gitea.io/sdk/gitea"
	"github.com/google/go-github/v66/github"
	log "github.com/sirupsen/logrus"
)

func CreateOrgFromUser (giteaClient *gitea.Client, user *github.User) ( created_user *gitea.Organization) {
	// The following block is horrible and I apologise to anyone reading this
	if Args.Rename != "" {
		user.Login = &Args.Rename
		// Does not feel right to do this, moral issue not a technical one
		// user.Name = &Args.Rename
	}

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
	if err != nil && strings.Contains(err.Error(), "user already exists") {
		err = nil
		log.Infof("User %s already exists", *user.Login)
	} else {
		CheckIfError(err)
	}
	err = ProcessAvatar(*user.Login, *user.AvatarURL)
	CheckIfError(err)

	return created_user
}

func CreateOrgFromOrg (giteaClient *gitea.Client, org *github.Organization) ( created_org *gitea.Organization) {
	// The following block is horrible and I apologise to anyone reading this
	if Args.Rename != "" {
		org.Login = &Args.Rename
	}

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
		log.Info("Importing " + *repo.Name)
		if repo.Description == nil {
			repo.Description = new(string)
			*repo.Description = ""
		}

		if !Args.DryRun {
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
			log.Info("Successfully imported " + *repo.Name)
		} else {
			log.Warn("Dry run: Would have imported " + *repo.Name)
		}
	}
	return err

}