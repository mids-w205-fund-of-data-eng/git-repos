// list command
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Org struct {
	Name      string
	repoNames []string
}

func NewOrg(orgName string) *Org {
	return &Org{orgName, []string{}}
}

func (o *Org) GetRepos(pattern string) *[]string {

	ctx := context.Background()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		Type:        "private",
		ListOptions: github.ListOptions{PerPage: 100},
	}
	repos, _, err := client.Repositories.ListByOrg(ctx, o.Name, opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(repos))
	for _, repo := range repos {
		repoName := repo.GetName()
		if strings.Contains(repoName, pattern) {
			o.repoNames = append(o.repoNames, repoName)
		}
	}

	return &o.repoNames
}

func (c *Org) Help() string {
	return "List org repos (detailed help information here)"
}
