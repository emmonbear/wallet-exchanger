package wallet

type WalletService interface{
	Deposit()
}

type service struct{}

func NewService() *service {
	return &service{}
}
