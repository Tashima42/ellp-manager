package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tashima42/ellp-manager/database"
)

func (cr *Controller) CreateEvent(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	grade := &database.Grade{}
	cr.Logger.Info(requestID, " unmarshal request body")
	if err := json.Unmarshal(c.Body(), grade); err != nil {
		return err
	}

	cr.Logger.Info(requestID, " validating body")
	if err := cr.Validate.Struct(grade); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	cr.Logger.Info(requestID, " starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	cr.Logger.Info(requestID, " creating grade")
	if err := database.CreateWorkshopGradeTxx(tx, grade); err != nil {
		cr.Logger.Error(err)
		return fmt.Errorf("failed to create grade %s", err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.JSON(map[string]bool{"success": true})
}
