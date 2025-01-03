package repository

import (
	"github.com/emmonbear/wallet-exchanger/internal/repository/auth"
	"github.com/emmonbear/wallet-exchanger/internal/repository/balance"
	"github.com/emmonbear/wallet-exchanger/internal/repository/exchange"
	"github.com/emmonbear/wallet-exchanger/internal/repository/postgres"
	"github.com/emmonbear/wallet-exchanger/internal/repository/wallet"
)

type Repository struct {
	auth.AuthRepository
	balance.BalanceRepository
	exchange.ExchangeRepository
	wallet.WalletRepository
}

func NewRepository(db postgres.Database) *Repository {
	return &Repository{
		AuthRepository:    postgres.NewAuthPostgres(db),
		BalanceRepository: postgres.NewBalancePostgres(db),
	}
}
