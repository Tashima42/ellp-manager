package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tashima42/ellp-manager/database"
)

func (cr *Controller) CreateGoal(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	goal := &database.Goal{}
	cr.Logger.Info(requestID, " unmarshal request body")
	if err := json.Unmarshal(c.Body(), goal); err != nil {
		return err
	}

	cr.Logger.Info(requestID, " validating body")
	if err := cr.Validate.Struct(goal); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	cr.Logger.Info(requestID, " starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	cr.Logger.Info(requestID, " creating goal")
	if err := database.CreateGoalTxx(tx, goal); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.JSON(map[string]bool{"success": true})
}

func (cr *Controller) CreateGoalAttachment(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	ga := &database.GoalAttachment{}
	cr.Logger.Info(requestID, " unmarshal request body")
	if err := json.Unmarshal(c.Body(), ga); err != nil {
		return err
	}

	cr.Logger.Info(requestID, " validating body")
	if err := cr.Validate.Struct(ga); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	cr.Logger.Info(requestID, " starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	cr.Logger.Info(requestID, " creating goal attachment")
	if err := database.CreateGoalAttachmentTxx(tx, ga); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.JSON(map[string]bool{"success": true})
}
