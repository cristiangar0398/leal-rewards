package services

import "time"

type RewardsService struct{}

func NewRewardsService() *RewardsService {
	return &RewardsService{}
}

func (s *RewardsService) IsPromotionPeriod(tradeID string, date time.Time) bool {
	switch tradeID {
	case "Sucursal1":
		return date.Month() == time.May && date.Day() >= 15 && date.Day() <= 30
	case "Sucursal2":
		return date.Month() == time.May && date.Day() >= 15 && date.Day() <= 20
	}
	return false
}

func (s *RewardsService) CalculatePointsAndCashback(amount float64, tradeID string, date time.Time) (int, float64) {
	points := int(amount / 1000)
	cashback := amount / 1000

	if s.IsPromotionPeriod(tradeID, date) {
		if tradeID == "Sucursal1" {
			points *= 2
			cashback *= 2
		} else if tradeID == "Sucursal2" && amount > 20000 {
			points = int(float64(points) * 1.3)
			cashback *= 1.3
		}
	}

	return points, cashback
}
