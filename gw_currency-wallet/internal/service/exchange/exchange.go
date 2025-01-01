package exchange

type ExchangeService interface{}

type service struct{}

func NewService() *service {
	return &service{}
}
