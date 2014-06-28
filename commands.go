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
	Action: actionByDict("common", dict.CommonDict),
}

var commandGreek = cli.Command{
	Name:  "greek",
	Usage: "greek mythological figures",
	Description: `
`,
	Action: actionByDict("greek", dict.GreekDict),
}

var commandNorse = cli.Command{
	Name:  "norse",
	Usage: "norse gods and goddesses",
	Description: `
`,
	Action: actionByDict("norse", dict.NorseDict),
}

var commandPerson = cli.Command{
	Name:  "person",
	Usage: "person name in British, French, Italy, Spain, Greek, Finalnd and Russia",
	Description: `
`,
	Action: actionByDict("person", dict.PersonDict),
}

func actionByDict(name string, dictFunc func() []string) func(*cli.Context) {
	return func(c *cli.Context) {
		var dict []string
		if isCached(name) {
			dict = cachedDict(name)
		} else {
			dict = dictFunc()
			cacheDict(name, dict)
		}

		for _, w := range dict {
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
