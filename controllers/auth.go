package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tashima42/ellp-manager/database"
	"github.com/tashima42/ellp-manager/hash"
)

type SignInUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (cr *Controller) SignIn(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	s := &SignInUser{}
	cr.Logger.Info(requestID, " unmarshal request body")
	if err := json.Unmarshal(c.Body(), s); err != nil {
		return err
	}

	cr.Logger.Info(requestID, " validate body")
	if err := cr.Validate.Struct(s); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	cr.Logger.Info(requestID, " starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	cr.Logger.Info(requestID, " looking for user with email "+s.Email)
	user, err := database.GetUserByEmailTxx(tx, s.Email)
	if err != nil {
		cr.Logger.Info(requestID, " error: "+err.Error())
		if !strings.Contains(err.Error(), "no rows in result set") {
			return err
		}
		return fiber.NewError(http.StatusNotFound, "email "+s.Email+" not found")
	}

	cr.Logger.Info(requestID, " checking password")
	if !hash.CheckPassword(user.Password, s.Password) {
		return fiber.NewError(http.StatusUnauthorized, "incorrect password")
	}

	jwt, err := hash.NewJWT(cr.JWTSecret, map[string]interface{}{
		"user": map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	})
	if err != nil {
		return errors.New("failed to generate jwt: " + err.Error())
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "auth-token"
	cookie.Value = jwt
	cookie.Expires = time.Now().Add(time.Hour * 24)
	c.Cookie(cookie)

	return c.JSON(map[string]interface{}{"token": jwt})
}
