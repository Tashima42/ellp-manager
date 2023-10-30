package database

import "database/sql"

type DocumentType string

const (
	DocumentTypeRG  UserRole = "rg"
	DocumentTypeCPF UserRole = "cpf"
)

type Document struct {
	ID        int       `db:"id" json"id"`
	Name             string       `json:"name"`
	UserID           string       `json:"userId"`
	RequesterID      string       `json:"requesterId"`
	Type             DocumentType `json:"type"`
	Address          string       `json:"address"`
	Accepted         bool         `json:"accepted"`
	RequesterComment string       `json:"requesterComment"`
}

func CreateDocument(db *sql.DB, document *Document) error {
	// result := db.Create(user)
	// return result.Error
	return nil
}
