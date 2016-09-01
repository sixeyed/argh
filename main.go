package main

import (
	"os"

	"github.com/gianarb/argh/cmd"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("argh", "0.0.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"serve": func() (cli.Command, error) {
			return &cmd.ServeCmd{}, nil
		},
	}

	exitStatus, _ := c.Run()

	os.Exit(exitStatus)
}