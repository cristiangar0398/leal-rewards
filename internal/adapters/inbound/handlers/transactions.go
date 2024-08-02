package handlers

type TransactionRequest struct {
	Document  string  `json:"document"`
	Amount    float64 `json:"amount"`
	TradeName string  `json:"trade_name"`
}

type TransactionResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}
