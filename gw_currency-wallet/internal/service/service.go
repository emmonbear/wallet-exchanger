package service

import "github.com/emmonbear/wallet-exchanger/internal/storage"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repo *storage.Storage) *Service {
	return &Service{}
}
