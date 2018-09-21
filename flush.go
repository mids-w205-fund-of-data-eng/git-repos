// flush command
package main

import (
	"github.com/mitchellh/cli"
)

type FlushCommand struct {
	Ui cli.Ui
}

func (c *FlushCommand) Run(_ []string) int {
	c.Ui.Output("Would flush org repositories here")
	return 0
}

func (c *FlushCommand) Help() string {
	return "flush org repos (detailed help information here)"
}

func (c *FlushCommand) Synopsis() string {
	return "flush org repos"
}
