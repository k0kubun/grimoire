package main

import (
	"github.com/codegangsta/cli"
	"github.com/k0kubun/grimoire/dict"
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
	Usage: "common English word list",
	Description: `
`,
	Action: dict.CommonDict,
}

var commandGreek = cli.Command{
	Name:  "greek",
	Usage: "greek mythological figures",
	Description: `
`,
	Action: dict.GreekDict,
}

var commandNorse = cli.Command{
	Name:  "norse",
	Usage: "norse gods and goddesses",
	Description: `
`,
	Action: dict.NorseDict,
}

var commandPerson = cli.Command{
	Name:  "person",
	Usage: "person name in British, French, Italy, Spain, Greek, Finalnd and Russia",
	Description: `
`,
	Action: dict.PersonDict,
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
