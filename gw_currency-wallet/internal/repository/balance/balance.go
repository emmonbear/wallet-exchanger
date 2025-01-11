package balance

import "github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/model"

type BalanceRepository interface {
	GetBalance(userID int) (model.Balance, error)
}
