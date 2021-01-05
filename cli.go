package main

import (
	"github.com/payfazz/tango/cli/command"
	util2 "github.com/payfazz/tango/util"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tango-cli"
	app.Usage = "Easy way to use tango"
	app.Author = "vekaputra & febrianram (fazzfinancial)"
	app.Version = "1.6.0"
	app.Commands = []cli.Command{
		command.InitCommand(),
		command.MakeCommand(),
		command.UpdateCommand(),
	}

	util2.PullRepoFolder()

	err := app.Run(os.Args)
	if nil != err {
		log.Fatal(err)
	}
}
