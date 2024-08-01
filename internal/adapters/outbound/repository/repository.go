package repository

import (
	"context"

	"github.com/cristiangar0398/leal-rewards/internal/models"
)

var (
	implementation Repository
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	Close() error
}

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func Close() error {
	return implementation.Close()
}
