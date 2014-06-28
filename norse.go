package main

import (
	"github.com/codegangsta/cli"
)

var commandNorse = cli.Command{
	Name:  "norse",
	Usage: "norse gods and goddesses",
	Description: `
`,
	Action: norseDict,
}

func norseDict(c *cli.Context) {
}
