package database

import (
	"database/sql"

	"github.com/cristiangar0398/leal-rewards/internal/models"
)

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
