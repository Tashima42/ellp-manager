package cmd

import (
	"github.com/tashima42/ellp-manager/database"
	"github.com/urfave/cli/v2"
)

func DBCommand() *cli.Command {
	return &cli.Command{
		Name:        "db",
		Usage:       "controls migrations and db related operations",
		Subcommands: []*cli.Command{migrateCommand()},
	}
}

func migrateCommand() *cli.Command {
	return &cli.Command{
		Name:        "migrate",
		Subcommands: []*cli.Command{migrateDownCommand()},
	}
}

func migrateDownCommand() *cli.Command {
	return &cli.Command{
		Name: "down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "db-path",
				Aliases:  []string{"p"},
				Usage:    "sqlite database path",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			db, err := database.Open(ctx.String("db-path"), true)
			if err != nil {
				return err
			}
			return database.Close(db)
		},
	}
}
