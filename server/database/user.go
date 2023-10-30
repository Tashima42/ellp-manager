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
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Address   string    `db:"address" json:"address"`
	Role      UserRole  `db:"role" json:"role"`
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
