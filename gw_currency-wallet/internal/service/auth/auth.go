package auth

import (
	"github.com/emmonbear/wallet-exchanger/internal/lib/security"
	"github.com/emmonbear/wallet-exchanger/internal/model"
	"github.com/emmonbear/wallet-exchanger/internal/repository/auth"
)

type AuthService interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type service struct {
	repo auth.AuthRepository
}

func NewService(repo auth.AuthRepository) *service {
	return &service{repo: repo}
}

func (s *service) CreateUser(user model.User) (int, error) {
	user.Password = security.GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *service) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, security.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	return security.GenerateJWTToken(user)
}

func (s *service) ParseToken(token string) (int, error) {
	return security.ParseJWTToken(token)
}
