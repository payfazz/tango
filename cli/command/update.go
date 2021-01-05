package command

import (
	"fmt"
	"github.com/payfazz/tango/util"

	"github.com/urfave/cli"
)

func updateScripts(gitBranch string) string {
	return `#!/bin/sh

set -x

cd $HOME/.tango

git reset --hard
git checkout ` + gitBranch + `
git pull
GO111MODULE=on go install
`
}

// MakeCommand defines cli.Command for make command
func UpdateCommand() cli.Command {
	var fileName = `tango-update.sh`

	return cli.Command{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "tango update [branch]",
		Action: func(c *cli.Context) {
			branch := c.Args().Get(0)
			if "" == branch {
				branch = "master"
			}

			fmt.Println("Updating tango..")
			err := util.RunScript(fileName, updateScripts(branch))
			if nil != err {
				panic(err)
			}
			fmt.Println("tango updated.")
		},
	}
}
