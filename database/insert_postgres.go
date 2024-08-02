package database

import (
	"context"

	"github.com/cristiangar0398/leal-rewards/internal/models"
)

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id , document , password) VALUES ($1, $2 ,$3)", user.Id, user.Cc, user.Password)
	return err
}

func (repo *PostgresRepository) InsertTrade(ctx context.Context, trade *models.Trade) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO trades (id , name , user_id , conversion_rate) VALUES ($1, $2 ,$3 , $4)", trade.Id, trade.TradeName, trade.UserID, trade.ConversionRate)
	return err
}

func (repo *PostgresRepository) InsertTransaction(ctx context.Context, transaction *models.Transaction) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO transactions (id, user_id, trade_id, amount) VALUES ($1, $2, $3, $4)", transaction.ID, transaction.UserID, transaction.TradeID, transaction.Amount)
	return err
}

func (repo *PostgresRepository) InsertRecordPoints(ctx context.Context, id string, userID string, tradeId string, points int) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO points ( id ,user_id, trade_id, points) VALUES ($1, $2, $3 ,$4)", id, userID, tradeId, points)
	return err
}

func (repo *PostgresRepository) InsertRecordCashback(ctx context.Context, id string, userID string, cashback float64) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO cashback (id , user_id, amount) VALUES ($1, $2 ,$3)", id, userID, cashback)
	return err
}
