package balance

import "github.com/emmonbear/wallet-exchanger/internal/model"

type BalanceRepository interface {
	GetBalance(userID int) (model.Balance, error)
}
