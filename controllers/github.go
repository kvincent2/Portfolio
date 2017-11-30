package controllers

import (
	"github.com/google/go-github/github"
	"context"
	"golang.org/x/oauth2"
	"os"
)


// GetHello controller
// add Github goSDK here
func GetProjects() ([]*github.Repository) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_PAT")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, _ := client.Repositories.List(ctx, "", nil)

	// for _, repo := range repos {
       
 //        fmt.Println(*repo.HTMLURL)
 //    }

	return repos
}

