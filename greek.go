package main

import (
	"github.com/codegangsta/cli"
)

var commandGreek = cli.Command{
	Name:  "greek",
	Usage: "greek mythological figures",
	Description: `
`,
	Action: greekDict,
}

func greekDict(c *cli.Context) {
}
