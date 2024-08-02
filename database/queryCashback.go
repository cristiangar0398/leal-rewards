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
