// list command
package main

import (
	"log"
)

type Repo struct {
	Name string

	gh        *GithubConnection
	repoNames []string
}

func NewRepo(repoName string) *Repo {

	gh, err := NewGithubConnection()
	if err != nil {
		log.Fatalf("GithubConnection: %s", err)
	}

	return &Repo{repoName, gh, []string{}}
}
