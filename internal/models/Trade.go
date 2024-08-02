package models

type Trade struct {
	Id             string `json:"id"`
	TradeName      string `json:"trade_name"`
	UserID         string `json:"user_id"`
	ConversionRate int    `json:"conversion_rate"` // Agregado el campo conversion_rate
}
