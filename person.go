package main

import (
	"github.com/codegangsta/cli"
)

var commandPerson = cli.Command{
	Name:  "person",
	Usage: "person name in British, French, Italy, Spain, Greek, Finalnd and Russia",
	Description: `
`,
	Action: personDict,
}

func personDict(c *cli.Context) {
}
