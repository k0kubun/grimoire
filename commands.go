package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandCommon,
	commandPerson,
	commandGreek,
	commandNorse,
}

var commandCommon = cli.Command{
	Name:  "common",
	Usage: "",
	Description: `
`,
	Action: doCommon,
}

var commandPerson = cli.Command{
	Name:  "person",
	Usage: "",
	Description: `
`,
	Action: doPerson,
}

var commandGreek = cli.Command{
	Name:  "greek",
	Usage: "",
	Description: `
`,
	Action: doGreek,
}

var commandNorse = cli.Command{
	Name:  "norse",
	Usage: "",
	Description: `
`,
	Action: doNorse,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doCommon(c *cli.Context) {
}

func doPerson(c *cli.Context) {
}

func doGreek(c *cli.Context) {
}

func doNorse(c *cli.Context) {
}
