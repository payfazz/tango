package main

import (
	"log"
	"os"

	"github.com/payfazz/tango/cli/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tango-cli"
	app.Usage = "Easy way to use tango"
	app.Author = "vekaputra & febrianram (fazzfinancial)"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		command.InitCommand(),
		command.MakeCommand(),
	}

	err := app.Run(os.Args)
	if nil != err {
		log.Fatal(err)
	}
}
