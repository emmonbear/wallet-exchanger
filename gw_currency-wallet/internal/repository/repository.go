package repository

import (
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository/auth"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository/balance"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository/exchange"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository/postgres"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository/wallet"
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
		WalletRepository:  postgres.NewWalletPostgres(db),
	}
}
