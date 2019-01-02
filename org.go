// list command
package main

import (
	"context"
	"log"
	"os"
	"regexp"

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

func (o *Org) GetReposMatching(pattern string) []string {

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
		matched, err := regexp.MatchString(pattern, repoName)
		if err == nil && matched {
			o.repoNames = append(o.repoNames, repoName)
		}
	}

	return o.repoNames
}

func (o *Org) CreateRepo(name string) (*github.Repository, error) {
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(true),
	}
	repo, _, err := o.client.Repositories.Create(o.ctx, o.Name, repo)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	return repo, err
}

func (o *Org) DeleteRepoByName(repoName string) error {
	_, err := o.client.Repositories.Delete(o.ctx, o.Name, repoName)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	return err
}

func (o *Org) DeleteReposByPattern(pattern string) error {
	log.Printf("deleting repos matching pattern: %s", pattern)
	repos := o.GetReposMatching(pattern)
	for _, repoName := range repos {
		log.Printf("removing repo %s", repoName)
		err := o.DeleteRepoByName(repoName)
		if err != nil {
			log.Printf("Error deleting repo %s: %s", repoName, err)
		}
	}
	return nil
}
