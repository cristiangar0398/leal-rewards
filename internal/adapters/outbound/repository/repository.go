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
	InsertTrade(ctx context.Context, trade *models.Trade) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByDocument(ctx context.Context, document string) (*models.User, error)
	GetTradeById(ctx context.Context, id string) (*models.Trade, error)
	GetUserDetail(ctx context.Context, document string) (*models.UserDetail, error)
	Close() error
}

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func InsertTrade(ctx context.Context, trade *models.Trade) error {
	return implementation.InsertTrade(ctx, trade)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByDocument(ctx context.Context, document string) (*models.User, error) {
	return implementation.GetUserByDocument(ctx, document)
}

func GetTradeById(ctx context.Context, id string) (*models.Trade, error) {
	return implementation.GetTradeById(ctx, id)
}

func GetUserDetail(ctx context.Context, document string) (*models.UserDetail, error) {
	return implementation.GetUserDetail(ctx, document)
}

func Close() error {
	return implementation.Close()
}
