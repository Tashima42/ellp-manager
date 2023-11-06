package main

import (
	"log"
	"os"

	"github.com/tashima42/ellp-manager/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:                   "ellp",
		Usage:                  "control and run the ellp manager",
		UseShortOptionHandling: true,
		Commands:               []*cli.Command{cmd.DBCommand(), cmd.ServerCommand()},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
