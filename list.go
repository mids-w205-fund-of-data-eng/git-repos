// list command
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/mitchellh/cli"
	"golang.org/x/oauth2"
)

type ListCommand struct {
	OrgName string
	Ui      cli.Ui
}

func (c *ListCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("list", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	cmdFlags.StringVar(&c.OrgName, "org-name", "", "The github org to list")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if c.OrgName == "" {
		log.Fatal("need to add an org to list")
	}

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
	repos, _, err := client.Repositories.ListByOrg(ctx, c.OrgName, opt)
	if err != nil {
		log.Fatal(err)
	}

	c.Ui.Output(fmt.Sprintf("Would list repositories for the org: %s", c.OrgName))
	c.Ui.Output(fmt.Sprintf("the org: %s has %d repos", c.OrgName, len(repos)))
	return 0
}

func (c *ListCommand) Help() string {
	return "List org repos (detailed help information here)"
}

func (c *ListCommand) Synopsis() string {
	return "List org repos"
}
