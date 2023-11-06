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
	ID              int64        `db:"id" json:"id"`
	Name            string       `db:"name" json:"name" validate:"required"`
	UserID          int64        `db:"user_id" json:"userId" validate:"required"`
	ReviewerID      int64        `db:"reviewer_id" json:"reviewerId" validate:"required"`
	Type            DocumentType `db:"type" json:"type" validate:"required"`
	Address         string       `db:"address" json:"address"`
	Accepted        bool         `db:"accepted" json:"accepted"`
	ReviewerComment string       `db:"reviewer_comment" json:"reviewerComment"`
	CreatedAt       time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time    `db:"updated_at" json:"updatedAt"`
}

func CreateDocumentTxx(tx *sqlx.Tx, d *Document) (id int64, err error) {
	query := "INSERT INTO documents(name, user_id, reviewer_id, type, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6);"
	result, err := tx.Exec(query, d.Name, d.UserID, d.ReviewerID, d.Type, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
