package services

import "time"

type RewardsServiceInterface interface {
	IsPromotionPeriod(tradeID string, date time.Time) bool
	CalculatePointsAndCashback(amount float64, tradeID string, date time.Time) (int, float64)
}
