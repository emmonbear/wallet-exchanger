package wallet

import "github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/model"

type WalletRepository interface {
	Deposit(wallet *model.Wallet) (float64, error)
	Withdraw(wallet *model.Wallet) (float64, error)
}
