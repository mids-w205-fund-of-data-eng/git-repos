package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	org = flag.String("org", "", "Github org")
)

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token")
	}
	if *org == "" {
		log.Fatal("need to add an org to list")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{Type: "private"}
	repos, _, err := client.Repositories.ListByOrg(ctx, *org, opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(repos))
}
