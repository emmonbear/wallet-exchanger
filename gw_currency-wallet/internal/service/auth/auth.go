package auth

type AuthService interface{}

type service struct{}

func NewService() *service {
	return &service{}
}
