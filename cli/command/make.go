package command

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/payfazz/tango/make/v1"
	make2 "github.com/payfazz/tango/make/v2"
	"github.com/payfazz/tango/make/v2/structure"
	"github.com/payfazz/tango/util"
	"io/ioutil"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var makeScripts = `#!/bin/sh

set -x

goimports -w ./internal
go fmt ./internal/...`

// MakeCommand defines cli.Command for make command
func MakeCommand() cli.Command {
	var fileName = `tango-make.sh`

	return cli.Command{
		Name:    "make",
		Aliases: []string{"m"},
		Usage:   "tango make <structure_path>",
		Action: func(c *cli.Context) {

			structureDefault := c.Args().Get(0)
			if structureDefault == "" {
				structureDefault = make.STRUCTURE_FILE
			}

			structurePath := ""
			_ = survey.AskOne(&survey.Input{
				Renderer: survey.Renderer{},
				Message:  "structure path:",
				Default:  structureDefault,
			}, &structurePath)

			if "" == structurePath {
				structurePath = make.STRUCTURE_FILE
			}

			// Read file and parse to struct
			content, err := ioutil.ReadFile(structurePath)
			if nil != err {
				panic(err)
			}

			var structureBase structure.Base
			err = yaml.Unmarshal(content, &structureBase)
			if nil != err {
				panic(err)
			}

			if structureBase.Version >= 2 {
				err = make2.Generate(structurePath, content, structureBase)
				if nil != err {
					panic(err)
				}
			} else {
				var structureMap make.StructureMap
				err = yaml.Unmarshal(content, &structureMap)
				if nil != err {
					panic(err)
				}

				// Generate stubs
				for _, strt := range structureMap.Structures {
					strt.Generate(make.DOMAIN_DIR)
				}
			}

			err = util.RunScript(fileName, makeScripts)
			if nil != err {
				panic(err)
			}
			fmt.Println("Structure generated.")
		},
	}
}
