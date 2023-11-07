package cmd

import (
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
	Port      int
	DB        *sqlx.DB
	JWTSecret []byte
	Logger    *zap.SugaredLogger
	Validate  *validator.Validate
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
			&cli.StringFlag{
				Name:     "jwt-secret",
				Usage:    "do not use this as a string flag, prefer setting the env var",
				Required: true,
				Aliases:  []string{"j"},
				EnvVars:  []string{"JWT_SECRET"},
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
		Port:      c.Int("port"),
		DB:        db,
		JWTSecret: []byte(c.String("jwt-secret")),
		Logger:    logger.Sugar(),
		Validate:  validator.New(validator.WithRequiredStructEnabled()),
	}
	return runServer(ec)
}

func runServer(ec *Context) error {
	cr := controllers.Controller{
		DB:        ec.DB,
		JWTSecret: ec.JWTSecret,
		Logger:    ec.Logger,
		Validate:  ec.Validate,
	}
	app := fiber.New(fiber.Config{ErrorHandler: cr.ErrorHandler})
	app.Use(requestid.New())
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("success")
	})
	app.Post("/signin", cr.SignIn)
	app.Use(cr.ValidateToken)
	app.Get("/hello", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*database.User)
		return c.SendString("Hello, " + user.Name)
	})
	app.Post("/user", cr.CreateUser)
	app.Post("/document", cr.CreateDocument)
	app.Post("/workshop", cr.CreateWorkshop)
	app.Post("/workshop/class", cr.CreateWorkshopClass)
	app.Post("/workshop/user", cr.CreateWorkshopUser)
	app.Post("/goal", cr.CreateGoal)
	app.Post("/goal/attachment", cr.CreateGoalAttachment)

	return app.Listen(":" + strconv.Itoa(ec.Port))
}
