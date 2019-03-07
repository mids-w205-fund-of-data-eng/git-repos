// list command
package main

import (
	"log"
	"regexp"

	"github.com/google/go-github/github"
)

type Org struct {
	Name string

	gh        *GithubConnection
	repoNames []string
}

func NewOrg(orgName string) *Org {
	gh, err := NewGithubConnection()
	if err != nil {
		log.Fatalf("GithubConnection: %s", err)
	}

	return &Org{orgName, gh, []string{}}
}

func (o *Org) GetReposMatching(pattern string) []string {

	options := &github.RepositoryListByOrgOptions{
		Type:        "private",
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var repos []*github.Repository
	for {
		page_of_repos, resp, err := o.gh.Client.Repositories.ListByOrg(o.gh.Context, o.Name, options)
		if err != nil {
			log.Fatal(err)
		}
		//log.Print("got a page of repos")
		repos = append(repos, page_of_repos...)
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	//log.Printf("got %d repos", len(repos))

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
	repo, _, err := o.gh.Client.Repositories.Create(o.gh.Context, o.Name, repo)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	return repo, err
}

func (o *Org) DeleteRepoByName(repoName string) error {
	_, err := o.gh.Client.Repositories.Delete(o.gh.Context, o.Name, repoName)
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
