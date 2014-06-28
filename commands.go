package main

import (
	"fmt"
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
	Action: func(c *cli.Context) {
		for _, w := range dict.CommonDict() {
			fmt.Println(w)
		}
	},
}

var commandGreek = cli.Command{
	Name:  "greek",
	Usage: "greek mythological figures",
	Description: `
`,
	Action: func(c *cli.Context) {
		for _, w := range dict.GreekDict() {
			fmt.Println(w)
		}
	},
}

var commandNorse = cli.Command{
	Name:  "norse",
	Usage: "norse gods and goddesses",
	Description: `
`,
	Action: func(c *cli.Context) {
		for _, w := range dict.NorseDict() {
			fmt.Println(w)
		}
	},
}

var commandPerson = cli.Command{
	Name:  "person",
	Usage: "person name in British, French, Italy, Spain, Greek, Finalnd and Russia",
	Description: `
`,
	Action: func(c *cli.Context) {
		for _, w := range dict.PersonDict() {
			fmt.Println(w)
		}
	},
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
