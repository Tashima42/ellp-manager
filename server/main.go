package main

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jmoiron/sqlx"
	"github.com/tashima42/ellp-manager/server/controllers"
	"github.com/tashima42/ellp-manager/server/database"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type Context struct {
	Port   int
	DB     *sqlx.DB
	Logger *zap.SugaredLogger
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

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	ec := &Context{
		Port:   c.Int("port"),
		DB:     db,
		Logger: logger.Sugar(),
	}
	return runServer(ec)
}

func runServer(ec *Context) error {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			err = ctx.Status(code).JSON(map[string]string{"error": err.Error()})
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			return nil
		},
	})
	app.Use(requestid.New())

	cr := controllers.Controller{DB: ec.DB, Logger: ec.Logger}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/user", cr.CreateUser)

	app.Listen(":" + strconv.Itoa(ec.Port))

	return nil
}
