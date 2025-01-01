package wallet

type WalletService interface{}
type service struct{}

func NewService() *service {
	return &service{}
}
