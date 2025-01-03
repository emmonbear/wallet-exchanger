package balance

import (
	"github.com/emmonbear/wallet-exchanger/internal/model"
	"github.com/emmonbear/wallet-exchanger/internal/repository/balance"
)

type BalanceService interface {
	GetBalance(userID int) (model.Balance, error)
}

type service struct {
	repo balance.BalanceRepository
}

func NewService(repo balance.BalanceRepository) *service {
	return &service{repo: repo}
}

func (s *service) GetBalance(userID int) (model.Balance, error) {
	return s.repo.GetBalance(userID)
}
