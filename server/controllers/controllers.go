package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/tashima42/ellp-manager/server/database"
	"github.com/tashima42/ellp-manager/server/hash"
)

type Controller struct {
	DB *sqlx.DB
}

func (cr *Controller) CreateUser(c *fiber.Ctx) error {
	user := &database.User{}
	if err := json.Unmarshal(c.Body(), user); err != nil {
		return err
	}
	hashedPassword, err := hash.Password(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	if _, err := database.GetUserByEmailTxx(tx, user.Email); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
    return c.Status(http.StatusConflict).JSON(map[string]string{"error": "email " + user.Email + " already was registered"})
	}

	if err := database.CreateUserTxx(tx, user); err != nil {
		return err
	}
	return c.JSON(map[string]interface{}{"success": true})
}
