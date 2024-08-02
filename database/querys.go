package database

import (
	"database/sql"

	"github.com/cristiangar0398/leal-rewards/internal/models"
)

func getCashbacksByUserId(db *sql.DB, userId string) ([]models.Cashback, error) {
	rows, err := db.Query("SELECT amount FROM cashback WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cashback []models.Cashback
	for rows.Next() {
		var cash models.Cashback
		if err := rows.Scan(&cash.Amount); err != nil {
			return nil, err
		}
		cashback = append(cashback, cash)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cashback, nil
}

func getPointsByUserId(db *sql.DB, userId string) ([]models.Point, error) {
	rows, err := db.Query("SELECT trade_id, SUM(points) FROM points WHERE user_id = $1 GROUP BY trade_id", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var points []models.Point
	for rows.Next() {
		var point models.Point
		if err := rows.Scan(&point.TradeId, &point.Points); err != nil {
			return nil, err
		}
		points = append(points, point)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

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
