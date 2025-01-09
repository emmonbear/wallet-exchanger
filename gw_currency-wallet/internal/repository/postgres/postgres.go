package postgres

import (
	"fmt"

	"github.com/emmonbear/wallet-exchanger/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	UsersTable        = "users"
	UserBalancesTable = "user_balances"
	CurrencyTable     = "currency"
)

type Storage struct {
	db *sqlx.DB
}

type Database interface {
	Close() error
	WithTransaction(fn TransactionFunc) error
}

type TransactionFunc func(tx *sqlx.Tx) error

func New(cfg *config.Config) (*Storage, error) {
	const op = "repository.postgres.New"

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.StorageConfig.DBHost,
		cfg.StorageConfig.DBPort,
		cfg.StorageConfig.DBUsername,
		cfg.StorageConfig.DBPassword,
		cfg.StorageConfig.DBName,
		cfg.StorageConfig.DBSSLMode,
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) WithTransaction(fn TransactionFunc) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction rollback failed: %w, original error: %w", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
