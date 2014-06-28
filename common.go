package main

import (
	"github.com/codegangsta/cli"
)

var commandCommon = cli.Command{
	Name:  "common",
	Usage: "common English word list",
	Description: `
`,
	Action: commonDict,
}

func commonDict(c *cli.Context) {
	println("common english")
}
