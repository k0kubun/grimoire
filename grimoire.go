package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func mainAction(c *cli.Context) {
	for _, w := range allDict() {
		fmt.Println(w)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "grimoire"
	app.Usage = "Naming utility for your software"
	app.Version = Version
	app.Commands = Commands
	app.Action = mainAction
	app.Run(os.Args)
}
