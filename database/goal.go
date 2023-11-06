package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Goal struct {
	ID            int64     `db:"id" json:"id"`
	CoordinatorID int64     `db:"coordinator_id" json:"coordinatorId" validate:"required"`
	Name          string    `db:"name" json:"name" validate:"required"`
	Description   string    `db:"description" json:"description"`
	Percentage    int       `db:"percentage" json:"percentage" validate:"required"`
	DueAt         time.Time `db:"due_at" json:"dueAt"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time `db:"updated_at" json:"updatedAt"`
}

type GoalAttachment struct {
	ID         int64     `db:"id" json:"id"`
	UserID     int64     `db:"user_id" json:"userId" validate:"required"`
	GoalID     int64     `db:"goal_id" json:"goalId" validate:"required"`
	Name       string    `db:"name" json:"name" validate:"required"`
	Comment    string    `db:"comment" json:"comment"`
	DocumentID int64     `db:"document_id" json:"documentId"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time `db:"updated_at" json:"updatedAt"`
}

func CreateGoalTxx(tx *sqlx.Tx, g *Goal) error {
	query := "INSERT INTO goals(coordinator_id, name, description, percentage, due_at, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7);"
	_, err := tx.Exec(query, g.CoordinatorID, g.Name, g.Description, g.Percentage, g.DueAt, time.Now(), time.Now())
	return err
}

func CreateGoalAttachmentTxx(tx *sqlx.Tx, ga *GoalAttachment) error {
	query := "INSERT INTO goal_attachments(user_id, goal_id, name, comment, document_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7);"
	_, err := tx.Exec(query, ga.UserID, ga.GoalID, ga.Name, ga.Comment, ga.DocumentID, time.Now(), time.Now())
	return err
}
