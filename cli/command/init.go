package command

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"runtime"

	"github.com/payfazz/tango/cli/util"
	"github.com/urfave/cli"
)

var missingProjectName = `Invalid project name, project name cannot be empty!`

var invalidProjectName = `Invalid project name, 'test' cannot be used for project name!`

var initLinuxScripts = `#!/bin/sh

set -x

mkdir $1
cp -a $HOME/.tango/template/default/. $1

cd $1
mv cmd/tango cmd/$1
find .ci cmd config database transport internal lib test -type f -exec sed -i'' "s/tango/$1/g" {} \;
find .ci cmd config database transport internal lib test -type f -exec sed -i'' "s/\/template\/default//g" {} \;
sed -i'' "s/tango\/template\/default/$1/g" go.mod
go mod tidy
rm -rf cli cli.go .git
cd ..`

var initMacScripts = `#!/bin/sh

set -x

mkdir $1
cp -a $HOME/.tango/template/default/. $1

cd $1
mv cmd/tango cmd/$1
find .ci cmd config database transport internal lib test -type f -exec sed -i '' "s/tango/$1/g" {} \;
find .ci cmd config database transport internal lib test -type f -exec sed -i '' "s/\/template\/default//g" {} \;
sed -i '' "s/tango\/template\/default/$1/g" go.mod
go mod tidy
cd ..`

// InitCommand defines cli.Command for init command
func InitCommand() cli.Command {
	var fileName = `tango-init.sh`

	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "tango init <project_name>",
		Action: func(c *cli.Context) {
			var err error

			projectName := ""
			_ = survey.AskOne(&survey.Input{
				Renderer: survey.Renderer{},
				Message:  "Project name:",
				Default:  c.Args().Get(0),
			}, &projectName)

			if "" == projectName {
				fmt.Println(missingProjectName)
				return
			}

			if "test" == projectName {
				fmt.Println(invalidProjectName)
				return
			}

			fmt.Println("Initialize project directory..")
			if runtime.GOOS == "darwin" { // mac
				err = util.RunScript(fileName, initMacScripts, projectName)
			} else { // linux
				err = util.RunScript(fileName, initLinuxScripts, projectName)
			}

			if nil != err {
				fmt.Println(err)
				return
			}
			fmt.Println("Project skeleton generated.")
		},
	}
}
