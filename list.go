// list command
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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
	cmdFlags.StringVar(&c.OrgName, "org-name", os.Getenv("GITHUB_ORG"), "The working GitHub Org")
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
		c.Ui.Output(fmt.Sprintf("%s", repoName))
	}

	return 0
}

func (c *ListCommand) Help() string {
	helpText := `
Usage: git repos list --org-name=<org_name> [<pattern>]
	Parameters:
	'<org_name>': The working GitHub Organization.  Default to the GITHUB_ORG environment variable.
	'<pattern>': An optional string to match repo names.
	`
	return strings.TrimSpace(helpText)
}

func (c *ListCommand) Synopsis() string {
	return "List org repos"
}
