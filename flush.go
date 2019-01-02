// flush command
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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
	cmdFlags.StringVar(&c.OrgName, "org-name", os.Getenv("GITHUB_ORG"), "The working GitHub Org")
	cmdFlags.BoolVar(&c.Confirmed, "confirm", false, "Really delete?")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	cmdArgs := cmdFlags.Args()
	if len(cmdArgs) >= 1 {
		c.Pattern = cmdArgs[0]
	}

	if c.OrgName == "" {
		c.Ui.Output(fmt.Sprint("Error: no GitHub Org"))
		cmdFlags.Usage()
		return 1
	}

	if c.Pattern == "" {
		c.Ui.Output(fmt.Sprint("Error: no filter pattern (not going to delete all org repos)"))
		cmdFlags.Usage()
		return 1
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
	helpText := `
Usage: git-repos flush --org-name=<org_name> --confirm <pattern>
	Parameters:
	'<org_name>': The working GitHub Organization.  Default to the GITHUB_ORG environment variable
	'confirm': Confirm repo deletion
	'<pattern>': A string to match repo names
	`
	return strings.TrimSpace(helpText)
}

func (c *FlushCommand) Synopsis() string {
	return "flush org repos"
}
