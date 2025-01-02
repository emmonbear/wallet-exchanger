package postgres

import (
	"fmt"

	"github.com/emmonbear/wallet-exchanger/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (username, pass_hash, email) VALUES ($1, $2, $3) RETURNING id", UserTable,
	)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Email)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND pass_hash=$2", UserTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
