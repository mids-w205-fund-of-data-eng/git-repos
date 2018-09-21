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

	var org Org
	o := org.NewOrg(c.OrgName)

	c.Ui.Output(fmt.Sprintf("Would list repositories for the org: %s", c.OrgName))
	c.Ui.Output(fmt.Sprintf("the org: %s has %d repos", c.OrgName, len(*o.GetRepos(""))))
	return 0
}

func (c *ListCommand) Help() string {
	return "List org repos (detailed help information here)"
}

func (c *ListCommand) Synopsis() string {
	return "List org repos"
}
