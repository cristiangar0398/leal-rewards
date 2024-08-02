package database

import (
	"database/sql"

	"github.com/cristiangar0398/leal-rewards/internal/models"
)

func getTradesData(db *sql.DB, query string, userId string) ([]models.Trade, error) {
	rows, err := db.Query(query, userId)
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
