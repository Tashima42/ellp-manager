package database

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var ErrInvalidEventType error = errors.New("invalid event type")

type EventType int

const (
	EventTypeMeeting EventType = 1
)

var EventTypes map[string]EventType = map[string]EventType{
	"meeting": EventTypeMeeting,
}

type Event struct {
	ID          int64     `db:"id" json:"id"`
	Type        EventType `db:"coordinator_id"`
	StringType  string    `json:"coordinatorId" validate:"required"`
	Description string    `db:"description" json:"description" validate:"required"`
	StartAt     time.Time `db:"start_at" json:"startAt"`
	EndAt       time.Time `db:"end_at" json:"endAt"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

type EventUser struct {
	ID           int64     `db:"id" json:"id"`
	EventID      int64     `db:"event_id" json:"eventId" validate:"required"`
	UserID       int64     `db:"user_id" json:"userId" validate:"required"`
	Participated bool      `db:"participated" json:"participated"`
	StartAt      time.Time `db:"start_at" json:"startAt"`
	EndAt        time.Time `db:"end_at" json:"endAt"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time `db:"updated_at" json:"updatedAt"`
}

func StringToEventType(key string) (EventType, error) {
	if e, ok := EventTypes[key]; ok {
		return e, nil
	}
	return -1, ErrInvalidEventType
}

func CreateEventTxx(tx *sqlx.Tx, e *Event) error {
	query := "INSERT INTO events(type, description, start_at, end_at, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6);"
	_, err := tx.Exec(query, e.Type, e.Description, e.StartAt, e.EndAt, time.Now(), time.Now())
	return err
}

func CreatEventUserTxx(tx *sqlx.Tx, eu *EventUser) error {
	query := "INSERT INTO event_user(event_id, user_id, created_at, updated_at) VALUES($1, $2, $3, $4);"
	_, err := tx.Exec(query, eu.EventID, eu.UserID, time.Now(), time.Now())
	return err
}
