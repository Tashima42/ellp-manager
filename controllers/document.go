package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tashima42/ellp-manager/database"
)

func (cr *Controller) CreateDocument(c *fiber.Ctx) error {
	requestID := fmt.Sprintf("%+v", c.Locals("requestid"))
	document := &database.Document{}
	cr.Logger.Info(requestID, " unmarshal request body")
	if err := json.Unmarshal(c.Body(), document); err != nil {
		return err
	}

	cr.Logger.Info(requestID, " validating body")
	if err := cr.Validate.Struct(document); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	cr.Logger.Info(requestID, " parse document type")
	dt, err := database.StringToDocumentType(document.StringType)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	document.Type = dt

	cr.Logger.Info(requestID, " starting transaction")
	tx, err := cr.DB.BeginTxx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	cr.Logger.Info(requestID, " looking for user with id ", document.UserID)
	if _, err := database.GetUserByIDTxx(tx, document.UserID); err != nil {
		cr.Logger.Info(requestID, " error: "+err.Error())
		if strings.Contains(err.Error(), "no rows in result set") {
			return fiber.NewError(http.StatusBadRequest, "user not found")
		}
		return err
	}

	cr.Logger.Info(requestID, " looking for reviewer user with id ", document.ReviewerID)
	if _, err := database.GetUserByIDTxx(tx, document.ReviewerID); err != nil {
		cr.Logger.Info(requestID, " error: "+err.Error())
		if strings.Contains(err.Error(), "no rows in result set") {
			return fiber.NewError(http.StatusBadRequest, "user not found")
		}
		return err
	}

	cr.Logger.Info(requestID, " creating document")
	documentID, err := database.CreateDocumentTxx(tx, document)
	if err != nil {
		return err
	}
	if err := database.CreateLogTxx(tx, &database.Log{
		Action:      "create",
		DocumentID:  documentID,
		UserID:      document.ReviewerID,
		Description: "created document " + document.Name,
	}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	cr.Logger.Info(requestID, " document created")
	return c.JSON(map[string]bool{"success": true})
}
