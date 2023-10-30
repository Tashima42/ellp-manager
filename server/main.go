package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/tashima42/ellp-manager/server/controllers"
	"github.com/tashima42/ellp-manager/server/database"
	"github.com/urfave/cli/v2"
)

type Context struct {
	Port int
	DB   *sqlx.DB
}

func main() {
	app := cli.App{
		Name:                   "ellp",
		Usage:                  "start the ellp-manager server",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "port",
				Usage:    "port number to bind the server",
				Required: true,
				Aliases:  []string{"p"},
				EnvVars:  []string{"PORT"},
			},
			&cli.StringFlag{
				Name:     "db-path",
				Usage:    "path for the sqlite database",
				Required: true,
				Aliases:  []string{"d"},
				EnvVars:  []string{"DB_PATH"},
			},
			&cli.BoolFlag{
				Name:     "migrate-down",
				Usage:    "migrate the database down and then up",
				Required: false,
				Aliases:  []string{"m"},
				EnvVars:  []string{"MIGRATE_DOWN"},
			},
		},
		Action: run,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	db, err := database.Open(c.String("db-path"), c.Bool("migrate-down"))
	if err != nil {
		return err
	}
	defer database.Close(db)

	ec := &Context{
		Port: c.Int("port"),
		DB:   db,
	}
	return runServer(ec)
}

func runServer(ec *Context) error {
	app := fiber.New()

	cr := controllers.Controller{DB: ec.DB}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/user", cr.CreateUser)

	app.Listen(":" + strconv.Itoa(ec.Port))

	return nil
}
