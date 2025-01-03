package postgres

import (
	"fmt"

	"github.com/emmonbear/wallet-exchanger/internal/model"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db Database
}

func NewBalancePostgres(db Database) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) GetBalance(userID int) (model.Balance, error) {
	var balance model.Balance

	err := r.db.WithTransaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("SELECT usd, rub, eur FROM %s WHERE user_id=$1", UserBalancesTable)
		err := tx.Get(&balance, query, userID)
		return err
	})

	return balance, err
}
