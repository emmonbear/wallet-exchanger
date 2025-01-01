package balance

type BalanceService interface{}

type service struct{}

func NewService() *service {
	return &service{}
}
