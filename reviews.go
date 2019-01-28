// reviews command
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/cli"
)

type ReviewsCommand struct {
	OrgName  string
	Reviewer string
	State    string
	Pattern  string
	Ui       cli.Ui
}

func (c *ReviewsCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("reviews", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&c.OrgName, "org-name", os.Getenv("GITHUB_ORG"), "The working GitHub Org")
	cmdFlags.StringVar(&c.Reviewer, "reviewer", "", "Reviews by this reviewer")
	cmdFlags.StringVar(&c.State, "state", "", "Reviews in this state")
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

	o := NewOrg(c.OrgName)
	matchingRepos := o.GetReposMatching(c.Pattern)
	for _, repoName := range matchingRepos {
		r := NewRepo(c.OrgName, repoName)
		c.Ui.Output(fmt.Sprintf("%s: %d", repoName, r.Grade()))
	}

	return 0
}

func (c *ReviewsCommand) Help() string {
	helpText := `
Usage: git repos reviews --org-name=<org_name> --reviewer=<reviewer> --state=<state> [<pattern>]
	Parameters:
	'<org_name>': The working GitHub Organization.  Default to the GITHUB_ORG environment variable.
	'<reviewr>': If set, show only reviews by this reviewer.  Default is not set.
	'<state>': If set, show only reviews in this state. Default is not set.
	'<pattern>': If set, show only reviews for repos matching pattern.  Default is not set.
	`
	return strings.TrimSpace(helpText)
}

func (c *ReviewsCommand) Synopsis() string {
	return "Show reviews for org repos matching conditions"
}
