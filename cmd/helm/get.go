package main

import (
	"errors"

	"github.com/codegangsta/cli"
	"github.com/deis/helm-dm/format"
)

func getCmd() cli.Command {
	return cli.Command{
		Name:   "get",
		Usage:  "Retrieves the supplied deployment",
		Action: func(c *cli.Context) { run(c, get) },
	}
}

func get(c *cli.Context) error {
	args := c.Args()
	if len(args) < 1 {
		return errors.New("First argument, deployment name, is required. Try 'helm get --help'")
	}
	name := args[0]
	deployment, err := client(c).GetDeployment(name)
	if err != nil {
		return err
	}
	return format.YAML(deployment)
}
