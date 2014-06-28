package main

import (
	"fmt"
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
	Usage: "common English word list",
	Description: `
`,
	Action: actionByDictName("common"),
}

var commandGreek = cli.Command{
	Name:  "greek",
	Usage: "greek mythological figures",
	Description: `
`,
	Action: actionByDictName("greek"),
}

var commandNorse = cli.Command{
	Name:  "norse",
	Usage: "list of people, items and places in Norse mythology",
	Description: `
`,
	Action: actionByDictName("norse"),
}

var commandPerson = cli.Command{
	Name:  "person",
	Usage: "person name in British, French, Italy, Spain, Greek, Finalnd and Russia",
	Description: `
`,
	Action: actionByDictName("person"),
}

func actionByDictName(name string) func(*cli.Context) {
	return func(c *cli.Context) {
		dict := loadDictByName(name)

		for _, w := range formatDict(dict) {
			fmt.Println(w)
		}
	}
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
