package database

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var ErrInvalidDocumentType error = errors.New("invalid document type")

type DocumentType int

const (
	DocumentTypeRG    DocumentType = 1
	DocumentTypeCPF   DocumentType = 2
	DocumentTypeOther DocumentType = 3
)

var DocumentTypes map[string]DocumentType = map[string]DocumentType{
	"rg":    DocumentTypeRG,
	"cpf":   DocumentTypeCPF,
	"other": DocumentTypeOther,
}

type Document struct {
	ID              int64        `db:"id" json:"id"`
	Name            string       `db:"name" json:"name" validate:"required"`
	UserID          int64        `db:"user_id" json:"userId" validate:"required"`
	ReviewerID      int64        `db:"reviewer_id" json:"reviewerId" validate:"required"`
	Type            DocumentType `db:"type"`
	StringType      string       `json:"type" validate:"required"`
	Address         string       `db:"address" json:"address"`
	Accepted        bool         `db:"accepted" json:"accepted"`
	ReviewerComment string       `db:"reviewer_comment" json:"reviewerComment"`
	CreatedAt       time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time    `db:"updated_at" json:"updatedAt"`
}

func StringToDocumentType(key string) (DocumentType, error) {
	if d, ok := DocumentTypes[key]; ok {
		return d, nil
	}
	return -1, ErrInvalidDocumentType
}

func CreateDocumentTxx(tx *sqlx.Tx, d *Document) (id int64, err error) {
	query := "INSERT INTO documents(name, user_id, reviewer_id, type, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6);"
	result, err := tx.Exec(query, d.Name, d.UserID, d.ReviewerID, d.Type, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
