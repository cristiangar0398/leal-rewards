package models

type Trade struct {
	Id        string `json:"id"`
	TradeName string `json:"trade_name"`
	UserID    string `json:"user_id"`
}
