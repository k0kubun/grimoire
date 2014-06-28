package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "grimoire"
	app.Version = Version
	app.Usage = "Naming utility for your software"
	app.Author = "Takashi Kokubun"
	app.Email = "takashikkbn@gmail.com"
	app.Commands = Commands
	app.Action = func(c *cli.Context) {
		println("all dict list")
	}

	app.Run(os.Args)
}
