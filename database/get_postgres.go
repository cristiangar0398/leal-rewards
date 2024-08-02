package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/cristiangar0398/leal-rewards/internal/models"
)

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	rows, err := repo.db.QueryContext(ctx, "SELECT id , document FROM users WHERE id = $1", id)

	if err != nil {
		log.Fatal(err)
		return &user, nil
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Cc); err != nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) GetUserByDocument(ctx context.Context, document string) (*models.User, error) {
	var user models.User
	rows, err := repo.db.QueryContext(ctx, "SELECT id, document , password FROM users WHERE document = $1", document)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatal(cerr)
		}
	}()

	if rows.Next() {
		if err := rows.Scan(&user.Id, &user.Cc, &user.Password); err != nil {
			return nil, err
		}
		return &user, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *PostgresRepository) GetUserDetail(ctx context.Context, document string) (*models.UserDetail, error) {
	var user models.User
	row := repo.db.QueryRowContext(ctx, "SELECT id , document FROM users WHERE document = $1", document)
	if err := row.Scan(&user.Id, &user.Cc); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	points, err := getPointsByUserId(repo.db, user.Id)
	if err != nil {
		return nil, err
	}

	cashback, err := getCashbacksByUserId(repo.db, user.Id)
	if err != nil {
		return nil, err
	}

	trades, err := getTradesData(repo.db, "SELECT id, name FROM trades WHERE user_id = $1", user.Id)
	if err != nil {
		return nil, err
	}

	userDetail := &models.UserDetail{
		Id:       user.Id,
		Cc:       user.Cc,
		Points:   points,
		Cashback: cashback,
		Trades:   trades,
	}

	return userDetail, nil

}

func (repo *PostgresRepository) GetTradeById(ctx context.Context, id string) (*models.Trade, error) {
	var trade models.Trade
	rows, err := repo.db.QueryContext(ctx, "SELECT id , name , user_id , conversion_rate FROM trades WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatal(cerr)
		}
	}()

	if rows.Next() {
		if err := rows.Scan(&trade.Id, &trade.TradeName, &trade.UserID, &trade.ConversionRate); err != nil {
			return nil, err
		}
		return &trade, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *PostgresRepository) GetTradeIDByName(ctx context.Context, name string) (*models.Trade, error) {
	var trade models.Trade
	rows, err := repo.db.QueryContext(ctx, "SELECT id ,name , user_id ,conversion_rate FROM trades WHERE name = $1", name)

	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatal(cerr)
		}
	}()

	if rows.Next() {
		if err := rows.Scan(&trade.Id, &trade.TradeName, &trade.UserID, &trade.ConversionRate); err != nil {
			return nil, err
		}
		return &trade, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}
