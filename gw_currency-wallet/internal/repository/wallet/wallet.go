package wallet

import "github.com/emmonbear/wallet-exchanger/internal/model"

type WalletRepository interface {
	Deposit(wallet *model.Wallet) (float64, error)
	Withdraw(wallet *model.Wallet) (float64, error)
}
