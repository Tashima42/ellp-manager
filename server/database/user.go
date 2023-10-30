package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRole string

const (
	UserRoleStudent     UserRole = "student"
	UserRoleCoordinator UserRole = "coordinator"
	UserRoleSecretary   UserRole = "secretary"
	UserRoleInstructor  UserRole = "instructor"
)

type User struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" validate:"required"`
	Email     string    `db:"email" json:"email" validate:"required"`
	Password  string    `db:"password" json:"password" validate:"required"`
	Address   string    `db:"address" json:"address" validate:"required"`
	Role      UserRole  `db:"role" json:"role" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

func CreateUserTxx(tx *sqlx.Tx, u *User) error {
	query := "INSERT INTO users(role, name, email, password, address, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7);"
	_, err := tx.Exec(query, u.Role, u.Name, u.Email, u.Password, u.Address, time.Now(), time.Now())
	return err
}

func GetUserByEmailTxx(tx *sqlx.Tx, email string) (*User, error) {
	var u User
	query := "SELECT id, name, email, password, address, role, created_at, updated_at FROM users WHERE email=$1 LIMIT 1;"
	err := tx.Get(&u, query, email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByIDTxx(tx *sqlx.Tx, id int64) (*User, error) {
	var u User
	query := "SELECT id, name, email, password, address, role, created_at, updated_at FROM users WHERE id=$1 LIMIT 1;"
	err := tx.Get(&u, query, id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
