package command

import (
	"fmt"

	"github.com/payfazz/tango/cli/util"

	"github.com/urfave/cli"
)

var missingProjectName = `Invalid use of 'init' command!
Please write your project name after 'init'
Usage: tango init my-new-project`

var invalidProjectName = `Invalid project name, 'test' cannot be used for project name!`

var initScripts = `#!/bin/sh

set -x

git clone git@github.com:payfazz/tango.git $1
cd $1
mv cmd/tango cmd/$1
find .ci cmd config database http internal lib test -type f -exec sed -i '' "s/tango/$1/g" {} \;
sed -i '' "s/tango/$1/g" go.mod
go mod tidy
rm -rf cli cli.go .git
cd ..`

// InitCommand defines cli.Command for init command
func InitCommand() cli.Command {
	var fileName = `tango-init.sh`

	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "tango init <project_name>",
		Action: func(c *cli.Context) {
			projectName := c.Args().Get(0)
			if "" == projectName {
				fmt.Println(missingProjectName)
				return
			}

			if "test" == projectName {
				fmt.Println(invalidProjectName)
				return
			}

			err := util.RunScript(fileName, initScripts, projectName)
			if nil != err {
				fmt.Println(err)
				return
			}
		},
	}
}
