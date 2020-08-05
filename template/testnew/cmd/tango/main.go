package main

import (
	"fmt"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/tango/template/default/config"
	"github.com/payfazz/tango/template/default/database/migration"
	"github.com/payfazz/tango/template/default/transport/container"
	grpcServer "github.com/payfazz/tango/template/default/transport/grpc/server"
	httpServer "github.com/payfazz/tango/template/default/transport/http/server"
	monitorServer "github.com/payfazz/tango/template/default/transport/monitor/server"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	app := cli.NewApp()
	app.Name = "tango"
	app.Usage = "tango service"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) {
		config.SetVerboseQuery()

		app := container.CreateAppContainer()

		grpc := grpcServer.CreateGrpcServer(app)
		grpc.Serve()

		monitor := monitorServer.CreateMonitorServer()
		monitor.Serve()

		http := httpServer.CreateHttpServer(app)
		http.Serve()
	}
	app.Commands = []cli.Command{
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "tango migrate",
			Action: func(c *cli.Context) {
				migrate()
			},
		},
		{
			Name:    "generate-sql",
			Aliases: []string{"g"},
			Usage:   "tango generate-sql",
			Action: func(c *cli.Context) {
				generateSQL()
			},
		},
		{
			Name:    "config",
			Aliases: []string{"g"},
			Usage:   "tango config",
			Action: func(c *cli.Context) {
				printConfig()
			},
		},
	}

	err := app.Run(os.Args)
	if nil != err {
		log.Fatal(err)
	}
}

func migrate() {
	config.SetVerboseQuery()
	fazzdb.Migrate(
		config.GetMigrateDb(),
		"tango-backend",
		config.ForceMigrate(),
		config.RunSeeder(),
		migration.Sequence...,
	)
}

const (
	SQL_DIR           = "./sql"
	SQL_FORMATTER_CLI = "sql-formatter-cli"
)

func generateSQL() {
	fmt.Println("generating sql migration files")

	if !isSQLFormatterCLIExist() {
		panic("sql-formatter-cli not found: please install sql-formatter-cli before continue; 'npm install -g sql-formatter-cli'")
	}

	queries := fazzdb.Raw(true, migration.Sequence...)

	err := os.MkdirAll(SQL_DIR, os.FileMode(0744))
	if nil != err {
		panic(err)
	}

	for i, v := range queries {
		generatedFile := fmt.Sprintf("%s/%d.sql", SQL_DIR, i+1)

		_, err = os.Stat(generatedFile)
		if !os.IsNotExist(err) {
			fmt.Println("file", generatedFile, "already exists, skip generating current migration")
			continue
		}

		err = ioutil.WriteFile(generatedFile, []byte(v), os.FileMode(0644))
		if nil != err {
			panic(err)
		}

		cmd := exec.Command(
			SQL_FORMATTER_CLI,
			"-i", generatedFile, "-o", generatedFile,
		)

		if err = cmd.Start(); nil != err {
			fmt.Println("ERROR: failed formatting file", generatedFile)
			continue
		}
	}

	fmt.Println("finish generate migration into sql files")
}

func isSQLFormatterCLIExist() bool {
	_, err := exec.LookPath(SQL_FORMATTER_CLI)
	return nil == err
}

func printConfig() {
	config.PrintEnv()
}
