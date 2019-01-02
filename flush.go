// flush command
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mitchellh/cli"
)

type FlushCommand struct {
	OrgName   string
	Pattern   string
	Confirmed bool
	Ui        cli.Ui
}

func (c *FlushCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("flush", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&c.OrgName, "org-name", "", "GitHub Org")
	cmdFlags.StringVar(&c.Pattern, "pattern", "", "Flush repos matching this pattern")
	cmdFlags.BoolVar(&c.Confirmed, "confirm", false, "Really delete?")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if c.OrgName == "" {
		log.Fatal("Please specify a GitHub Org")
	}

	if c.Pattern == "" {
		log.Fatal("Please specify a pattern.  Not going to delete all repos in the org")
	}

	if c.Confirmed {
		c.Ui.Output(fmt.Sprintf("deleting the following repos:"))
	} else {
		c.Ui.Output(fmt.Sprintf("please '--confirm' deletion of the following repos:"))
	}

	o := NewOrg(c.OrgName)
	matchingRepos := o.GetReposMatching(c.Pattern)
	for _, repoName := range matchingRepos {
		c.Ui.Output(fmt.Sprintf("%s", repoName))
		if c.Confirmed {
			o.DeleteRepoByName(repoName)
		}
	}

	return 0
}

func (c *FlushCommand) Help() string {
	return "flush org repos (detailed help information here)"
}

func (c *FlushCommand) Synopsis() string {
	return "flush org repos"
}
