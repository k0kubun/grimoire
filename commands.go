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
