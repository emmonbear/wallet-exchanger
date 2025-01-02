package postgres

import (
	"fmt"

	"github.com/emmonbear/wallet-exchanger/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	UserTable     = "users"
	WalletTable   = "wallets"
	CurrencyTable = "currency"
)

type Storage struct {
	db *sqlx.DB
}

func New(cfg *config.Config) (*sqlx.DB, error) {
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
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
