package database

import (
	"database/sql"

	"github.com/cristiangar0398/leal-rewards/internal/models"
)

func getTradesByUserId(db *sql.DB, userId string) ([]models.Trade, error) {
	rows, err := db.Query("SELECT id, name FROM trades WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades []models.Trade
	for rows.Next() {
		var trade models.Trade
		if err := rows.Scan(&trade.Id, &trade.TradeName); err != nil {
			return nil, err
		}
		trades = append(trades, trade)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trades, nil
}
