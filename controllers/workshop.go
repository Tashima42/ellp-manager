package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tashima42/ellp-manager/database"
)

func (cr *Controller) CreateWorkshop(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	workshop := &database.Workshop{}
	cr.Logger.Info(requestID, " unmarshal request body")
	if err := json.Unmarshal(c.Body(), workshop); err != nil {
		return err
	}

	cr.Logger.Info(requestID, " validating body")
	if err := cr.Validate.Struct(workshop); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	cr.Logger.Info(requestID, " starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	cr.Logger.Info(requestID, " creating workshop")
	if err := database.CreateWorkshopTxx(tx, workshop); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.JSON(map[string]bool{"success": true})
}

func (cr *Controller) CreateWorkshopClass(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	workshopClass := &database.WorkshopClass{}
	cr.Logger.Info(requestID, " unmarshal request body")
	if err := json.Unmarshal(c.Body(), workshopClass); err != nil {
		return err
	}

	cr.Logger.Info(requestID, " validating body")
	if err := cr.Validate.Struct(workshopClass); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	cr.Logger.Info(requestID, " starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	cr.Logger.Info(requestID, " creating workshop class")
	if err := database.CreateWorkshopClassTxx(tx, workshopClass); err != nil {
		cr.Logger.Error(err)
		return fmt.Errorf("failed to create workshop class %s", err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.JSON(map[string]bool{"success": true})
}

func (cr *Controller) CreateWorkshopUser(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	workshopUser := &database.WorkshopUser{}
	cr.Logger.Info(requestID, " unmarshal request body")
	if err := json.Unmarshal(c.Body(), workshopUser); err != nil {
		return err
	}

	cr.Logger.Info(requestID, " validating body")
	if err := cr.Validate.Struct(workshopUser); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	cr.Logger.Info(requestID, " parse workshop user role")
	r, err := database.StringToWorkshopUserRole(workshopUser.StringRole)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	workshopUser.Role = r

	cr.Logger.Info(requestID, " starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	cr.Logger.Info(requestID, " creating workshop user")
	if err := database.CreateWorkshopUserTxx(tx, workshopUser); err != nil {
		cr.Logger.Error(err)
		return fmt.Errorf("failed to create workshop user %s", err.Error())
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.JSON(map[string]bool{"success": true})
}

func (cr *Controller) CreateGrade(c *fiber.Ctx) error {
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
