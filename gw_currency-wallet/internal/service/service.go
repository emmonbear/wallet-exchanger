package service

import (
	"github.com/emmonbear/wallet-exchanger/internal/repository"
	"github.com/emmonbear/wallet-exchanger/internal/service/auth"
	"github.com/emmonbear/wallet-exchanger/internal/service/balance"
	"github.com/emmonbear/wallet-exchanger/internal/service/exchange"
	"github.com/emmonbear/wallet-exchanger/internal/service/wallet"
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
		BalanceService:  balance.NewService(),
		ExchangeService: exchange.NewService(),
		WalletService:   wallet.NewService(),
	}
}
