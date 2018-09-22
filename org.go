// list command
package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Org struct {
	Name      string
	repoNames []string
	ctx       context.Context
	client    *github.Client
}

func NewOrg(orgName string) *Org {

	ctx := context.Background()

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &Org{orgName, []string{}, ctx, client}
}

func (o *Org) GetRepos(pattern string) []string {

	options := &github.RepositoryListByOrgOptions{
		Type:        "private",
		ListOptions: github.ListOptions{PerPage: 100},
	}
	repos, _, err := o.client.Repositories.ListByOrg(o.ctx, o.Name, options)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		repoName := repo.GetName()
		if strings.Contains(repoName, pattern) {
			o.repoNames = append(o.repoNames, repoName)
		}
	}

	return o.repoNames
}
