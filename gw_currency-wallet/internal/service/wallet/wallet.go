package wallet

import (
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/model"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository/wallet"
)

type WalletService interface {
	Deposit(wallet *model.Wallet) (float64, error)
	Withdraw(wallet *model.Wallet) (float64, error)
}

type service struct {
	repo wallet.WalletRepository
}

func NewService(repo wallet.WalletRepository) *service {
	return &service{repo: repo}
}

func (s *service) Deposit(wallet *model.Wallet) (float64, error) {
	return s.repo.Deposit(wallet)
}

func (s *service) Withdraw(wallet *model.Wallet) (float64, error) {
	return s.repo.Withdraw(wallet)
}
