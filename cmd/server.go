package cmd

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jmoiron/sqlx"
	"github.com/tashima42/ellp-manager/controllers"
	"github.com/tashima42/ellp-manager/database"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type Context struct {
	Port     int
	DB       *sqlx.DB
	Logger   *zap.SugaredLogger
	Validate *validator.Validate
}

type GlobalError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "start the ellp-manager server",
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
		},
		Action: run,
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
		Port:     c.Int("port"),
		DB:       db,
		Logger:   logger.Sugar(),
		Validate: validator.New(validator.WithRequiredStructEnabled()),
	}
	return runServer(ec)
}

func runServer(ec *Context) error {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			zap.Error(err)
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			err = ctx.Status(code).JSON(GlobalError{Success: false, Message: e.Error()})
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			return nil
		},
	})
	app.Use(requestid.New())

	cr := controllers.Controller{
		DB:       ec.DB,
		Logger:   ec.Logger,
		Validate: ec.Validate,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/user", cr.CreateUser)
	app.Post("/document", cr.CreateDocument)
	app.Post("/workshop", cr.CreateWorkshop)

	app.Listen(":" + strconv.Itoa(ec.Port))

	return nil
}
