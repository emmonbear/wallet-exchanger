package auth

import (
	"crypto/sha1"
	"fmt"

	"github.com/emmonbear/wallet-exchanger/internal/model"
	"github.com/emmonbear/wallet-exchanger/internal/repository/auth"
)

const salt = "1asdqwe2@#!@sadwqe#!,;dgl"

type AuthService interface {
	CreateUser(user model.User) (int, error)
}

type service struct {
	repo auth.AuthRepository
}

func NewService(repo auth.AuthRepository) *service {
	return &service{repo: repo}
}

func (s *service) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
