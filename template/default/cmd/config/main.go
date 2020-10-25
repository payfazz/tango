package main

import (
	"github.com/payfazz/tango/template/default/config"
	"log"
	"os"
)

func main() {
	if err := config.PrintEnv(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
