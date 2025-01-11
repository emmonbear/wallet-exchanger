package postgres

import (
	"fmt"

	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/model"
	"github.com/jmoiron/sqlx"
)

type WalletPostgres struct {
	db Database
}

func NewWalletPostgres(db Database) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) Deposit(wallet *model.Wallet) (float64, error) {
	var column string
	switch wallet.Currency {
	case "USD":
		column = "usd"
	case "RUB":
		column = "rub"
	case "EUR":
		column = "eur"
	default:
		return 0.0, fmt.Errorf("unsupported currency: %s", wallet.Currency)
	}

	var newBalance float64
	err := r.db.WithTransaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf(`
			UPDATE %s
			SET %s = %s + $1
			WHERE user_id = $2
			RETURNING %s
		`, UserBalancesTable, column, column, column)

		err := tx.QueryRow(query, wallet.Amount, wallet.UserID).Scan(&newBalance)
		if err != nil {
			return fmt.Errorf("failed to execute query: %w", err)
		}
		return nil
	})
	if err != nil {
		return 0.0, fmt.Errorf("transaction failed: %w", err)
	}

	return newBalance, nil
}

func (r *WalletPostgres) Withdraw(wallet *model.Wallet) (float64, error) {
	var column string
	switch wallet.Currency {
	case "USD":
		column = "usd"
	case "RUB":
		column = "rub"
	case "EUR":
		column = "eur"
	default:
		return 0.0, fmt.Errorf("unsupported currency: %s", wallet.Currency)
	}

	var newBalance float64

	err := r.db.WithTransaction(func(tx *sqlx.Tx) error {
		var currentBalance float64
		checkQuery := fmt.Sprintf(`
			SELECT %s
			FROM %s
			WHERE user_id = $1
		`, column, UserBalancesTable)

		err := tx.QueryRow(checkQuery, wallet.UserID).Scan(&currentBalance)
		if err != nil {
			return fmt.Errorf("failed to retrieve current balance: %w", err)
		}

		if currentBalance < wallet.Amount {
			return fmt.Errorf(
				"insufficient funds: current balance %.2f, requested amount %.2f",
				currentBalance, wallet.Amount,
			)
		}

		query := fmt.Sprintf(`
			UPDATE %s
			SET %s = %s - $1
			WHERE user_id = $2
			RETURNING %s
		`, UserBalancesTable, column, column, column)

		err = tx.QueryRow(query, wallet.Amount, wallet.UserID).Scan(&newBalance)
		if err != nil {
			return fmt.Errorf("failed to execute update query: %w", err)
		}

		return nil
	})

	if err != nil {
		return 0.0, fmt.Errorf("transaction failed: %w", err)
	}

	return newBalance, nil
}
