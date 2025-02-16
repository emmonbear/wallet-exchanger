package service

import (
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/service/auth"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/service/balance"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/service/exchange"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/service/wallet"
)

type Service struct {
	auth.AuthService
	balance.BalanceService
	exchange.ExchangeService
	wallet.WalletService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthService:     auth.NewService(repos),
		BalanceService:  balance.NewService(repos),
		ExchangeService: exchange.NewService(),
		WalletService:   wallet.NewService(repos),
	}
}
