package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
	c := cli.NewCLI("git-org", "0.0.3")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &ListCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
		"flush": func() (cli.Command, error) {
			return &FlushCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
