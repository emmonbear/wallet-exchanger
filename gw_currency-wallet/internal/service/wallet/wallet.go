package wallet

import (
	"github.com/emmonbear/wallet-exchanger/internal/model"
	"github.com/emmonbear/wallet-exchanger/internal/repository/wallet"
)

type WalletService interface {
	Deposit(wallet *model.Wallet) (float64, error)
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
