package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Workshop struct {
	ID            int64     `db:"id" json:"id"`
	CoordinatorID int64     `db:"coordinator_id" json:"coordinatorId" validate:"required"`
	Name          string    `db:"name" json:"name" validate:"required"`
	Description   string    `db:"description" json:"description" validate:"required"`
	Code          string    `db:"code" json:"code" validate:"required"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time `db:"updated_at" json:"updatedAt"`
}

func CreateWorkshopTxx(tx *sqlx.Tx, w *Workshop) error {
	query := "INSERT INTO workshops(coordinator_id, name, description, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6);"
	_, err := tx.Exec(query, w.CoordinatorID, w.Name, w.Description, w.Code, time.Now(), time.Now())
	return err
}
