package postgres

import (
	"fmt"

	"github.com/emmonbear/wallet-exchanger/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db Database
}

func NewAuthPostgres(db Database) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// BUG Индекс обновляется несмотря на rollback транзакции
func (r *AuthPostgres) CreateUser(user model.User) error {
	return r.db.WithTransaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf(
			"INSERT INTO %s (username, pass_hash, email) VALUES ($1, $2, $3) RETURNING id", UsersTable,
		)
		row := tx.QueryRow(query, user.Username, user.Password, user.Email)
		var id int
		if err := row.Scan(&id); err != nil {
			return err
		}

		balanceQuery := fmt.Sprintf("INSERT INTO %s (user_id) VALUES ($1)", UserBalancesTable)
		if _, err := tx.Exec(balanceQuery, id); err != nil {
			return err
		}

		return nil
	})

}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User

	err := r.db.WithTransaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND pass_hash=$2", UsersTable)
		err := tx.Get(&user, query, username, password)
		return err
	})

	return user, err
}
