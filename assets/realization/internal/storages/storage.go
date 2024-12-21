package storage

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, report interface{}) error
	DeleteAll(ctx context.Context) error

	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
