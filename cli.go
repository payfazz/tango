package main

import (
	"github.com/payfazz/tango/cli/command"
	"github.com/payfazz/tango/cli/util"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tango-cli"
	app.Usage = "Easy way to use tango"
	app.Author = "vekaputra & febrianram (fazzfinancial)"
	app.Version = "2.0.0"
	app.Commands = []cli.Command{
		command.InitCommand(),
		command.MakeCommand(),
	}

	const GitBranch = "new"
	util.PullRepoFolder(GitBranch)

	err := app.Run(os.Args)
	if nil != err {
		log.Fatal(err)
	}
}
