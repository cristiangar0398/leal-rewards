package models

type Transaction struct {
	ID      string  `json:"id"`
	UserID  string  `json:"user_id"`
	TradeID string  `json:"trade_id"`
	Amount  float64 `json:"amount"`
}
