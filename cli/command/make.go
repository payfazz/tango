package command

import (
	"io/ioutil"
	"os"

	"github.com/payfazz/tango/cli/util"
	"github.com/payfazz/tango/cli/util/make"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var makeScripts = `#!/bin/sh

set -x

goimports -w ./internal`

// MakeCommand defines cli.Command for make command
func MakeCommand() cli.Command {
	var fileName = `tango-make.sh`

	return cli.Command{
		Name:    "make",
		Aliases: []string{"m"},
		Usage:   "tango make <path_to_structure>",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "force, f",
				Usage: "force generate even if file already exist",
			},
		},
		Action: func(c *cli.Context) {
			structurePath := c.Args().Get(0)
			if "" == structurePath {
				structurePath = "./make/structure.yaml"
			}

			// Read file and parse to struct
			content, err := ioutil.ReadFile(structurePath)
			if nil != err {
				panic(err)
			}

			var structureMap make.StructureMap
			err = yaml.Unmarshal(content, &structureMap)
			if nil != err {
				panic(err)
			}

			// Move domain to backup dir
			dir := "./internal/domain"
			backupDir := "./internal/old"

			if c.Bool("force") {
				_ = os.Rename(dir, backupDir)
			}

			// Generate stubs
			for _, structure := range structureMap.Structures {
				make.GenerateStubs(structure)
			}

			err = util.RunScript(fileName, makeScripts)
			if nil != err {
				panic(err)
			}
		},
	}
}
