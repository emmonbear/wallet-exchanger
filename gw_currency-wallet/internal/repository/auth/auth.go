package auth

import "github.com/emmonbear/wallet-exchanger/internal/model"

type AuthRepository interface {
	CreateUser(user model.User) (int, error)
}
