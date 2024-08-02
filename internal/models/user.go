package models

type User struct {
	Id       string `json:"id"`
	Cc       string `json:"document"`
	Password string `json:"password"`
}

type UserDetail struct {
	Id       string     `json:"id"`
	Cc       string     `json:"document"`
	Points   []Point    `json:"points"`
	Cashback []Cashback `json:"cashback"`
	Trades   []Trade    `json:"trades"`
}

type Point struct {
	TradeId string `json:"trade_id"`
	Points  int    `json:"points"`
}

type Cashback struct {
	Amount float64 `json:"amount"`
}
