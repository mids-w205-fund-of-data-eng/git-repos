// list command
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mitchellh/cli"
)

type ListCommand struct {
	OrgName string
	Pattern string
	Ui      cli.Ui
}

func (c *ListCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("list", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&c.OrgName, "org-name", "", "The github org to list")
	cmdFlags.StringVar(&c.Pattern, "pattern", "", "List repos matching this pattern")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if c.OrgName == "" {
		log.Fatal("need to add an org to list")
	}

	o := NewOrg(c.OrgName)
	matchingRepos := *o.GetRepos(c.Pattern)

	c.Ui.Output(fmt.Sprintf("Listing repositories for %s containing %s", c.OrgName, c.Pattern))
	c.Ui.Output(fmt.Sprintf("the org: %s has %d repos that contain %s", c.OrgName, len(matchingRepos), c.Pattern))

	for _, repoName := range matchingRepos {
		c.Ui.Output(fmt.Sprintf("%s", repoName))
	}

	return 0
}

func (c *ListCommand) Help() string {
	return "List org repos (detailed help information here)"
}

func (c *ListCommand) Synopsis() string {
	return "List org repos"
}
