package command

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/payfazz/tango/cli/util"
	"github.com/urfave/cli"
)

var missingProjectName = `Invalid use of 'init' command!
Please write your project name after 'init'
Usage: tango init my-new-project`

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
sed -i'' "s/tango\/template\/default/$1/g" go.mod
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

			projectName := c.Args().Get(0)
			if "" == projectName {
				fmt.Println(missingProjectName)
				return
			}

			if "test" == projectName {
				fmt.Println(invalidProjectName)
				return
			}

			homeDir, _ := os.UserHomeDir()
			tangoDir := homeDir + "/.tango"
			if _, err = os.Stat(homeDir + "/.tango"); os.IsNotExist(err) {
				cmd := exec.Command("git", "clone", "git@github.com:payfazz/tango.git", tangoDir)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					fmt.Println(err)
					return
				}

				cmd = exec.Command("sh", "-c", "cd "+tangoDir+" && git checkout new && git pull")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			if runtime.GOOS == "darwin" { // mac
				err = util.RunScript(fileName, initMacScripts, projectName)
			} else { // linux
				err = util.RunScript(fileName, initLinuxScripts, projectName)
			}

			if nil != err {
				fmt.Println(err)
				return
			}
		},
	}
}
