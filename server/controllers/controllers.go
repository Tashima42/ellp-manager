package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/tashima42/ellp-manager/server/database"
	"github.com/tashima42/ellp-manager/server/hash"
)

type Controller struct {
	DB *sqlx.DB
}

func (cr *Controller) CreateUser(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	log.Println(requestID + " - create user")
	user := &database.User{}
	log.Println(requestID + " - unmarshal request body")
	if err := json.Unmarshal(c.Body(), user); err != nil {
		return err
	}

	log.Println(requestID + " - starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	log.Println(requestID + " - looking for user with email " + user.Email)
	if _, err := database.GetUserByEmailTxx(tx, user.Email); err != nil {
		log.Println(requestID + " - error: " + err.Error())
		if !strings.Contains(err.Error(), "no rows in result set") {
			return err
		}
		log.Println(requestID + " - user doesn't exists, continue")
	} else {
		log.Println(requestID + " - error: email was already registered")
		return c.Status(http.StatusConflict).JSON(map[string]string{"error": "email " + user.Email + " already was registered"})
	}

	log.Println(requestID + " - hashing password")
	hashedPassword, err := hash.Password(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	log.Println(requestID + " - creating user")
	if err := database.CreateUserTxx(tx, user); err != nil {
		return err
	}
	log.Println(requestID + " - user created")
	return c.JSON(map[string]interface{}{"success": true})
}
