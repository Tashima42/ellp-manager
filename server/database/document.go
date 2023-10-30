package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type DocumentType string

const (
	DocumentTypeRG  UserRole = "rg"
	DocumentTypeCPF UserRole = "cpf"
)

type Document struct {
	ID               int          `db:"id" json:"id"`
	Name             string       `db:"name" json:"name"`
	UserID           string       `db:"user_id" json:"userId"`
	RequesterID      string       `db:"requester_id" json:"requesterId"`
	Type             DocumentType `db:"type" json:"type"`
	Address          string       `db:"address" json:"address"`
	Accepted         bool         `db:"accepted" json:"accepted"`
	RequesterComment string       `db:"requester_comment" json:"requesterComment"`
	CreatedAt        time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time    `db:"updated_at" json:"updatedAt"`
}

func CreateDocumentTxx(tx *sqlx.DB, d *Document) error {
	query := "INSERT INTO documents(name, user_id, requester_id, type, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6);"
	_, err := tx.Exec(query, d.Name, d.UserID, d.RequesterID, d.Type, time.Now(), time.Now())
	return err
}
