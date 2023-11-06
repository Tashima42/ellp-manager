package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type LogAction string

const (
	LogActionCreate LogAction = "create"
	LogActionUpdate LogAction = "update"
	LogActionDelete LogAction = "delete"
	LogActionRead   LogAction = "read"
)

type Log struct {
	ID          int64     `db:"id" json:"id"`
	Action      string    `db:"action" json:"action" validate:"required"`
	DocumentID  int64     `db:"document_id" json:"documentId"`
	UserID      int64     `db:"user_id" json:"userId"`
	Description string    `db:"description" json:"description" validate:"required"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

func CreateLogTxx(tx *sqlx.Tx, l *Log) error {
	query := "INSERT INTO logs(action, user_id, document_id, description, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6);"
	_, err := tx.Exec(query, l.Action, l.UserID, l.DocumentID, l.Description, time.Now(), time.Now())
	return err
}
