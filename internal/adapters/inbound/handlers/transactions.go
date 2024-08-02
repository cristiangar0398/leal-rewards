package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/cristiangar0398/leal-rewards/internal/adapters/outbound/repository"
	"github.com/cristiangar0398/leal-rewards/internal/domain/services"
	"github.com/cristiangar0398/leal-rewards/internal/models"
	"github.com/cristiangar0398/leal-rewards/internal/server"
	"github.com/segmentio/ksuid"
)

type TransactionRequest struct {
	Document  string  `json:"document"`
	Amount    float64 `json:"amount"`
	TradeName string  `json:"trade_name"`
}

type TransactionResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func TransactionProcessHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := services.NewRewardsService()
		request, err := decodeTransactionRequest(r)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := repository.GetUserByDocument(r.Context(), request.Document)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		trade, err := repository.GetTradeIDByName(r.Context(), request.TradeName)
		if err != nil {
			http.Error(w, "Trade not found", http.StatusNotFound)
			return
		}

		transactionID, err := generateRandomID()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		transaction := &models.Transaction{
			ID:      transactionID,
			UserID:  user.Id,
			TradeID: trade.Id,
			Amount:  request.Amount,
		}

		err = repository.InsertTransaction(r.Context(), transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		date := time.Now()
		points, cashback := service.CalculatePointsAndCashback(request.Amount, trade.Id, date)

		pointsID, _ := generateRandomID()
		err = repository.InsertRecordPoints(r.Context(), pointsID, user.Id, trade.Id, points)
		if err != nil {
			log.Printf("Error recording points: %v", err)
			http.Error(w, "Error recording points", http.StatusInternalServerError)
			return
		}

		cashbackID, _ := generateRandomID()
		err = repository.InsertRecordCashback(r.Context(), cashbackID, user.Id, cashback)
		if err != nil {
			http.Error(w, "Error recording cashback", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TransactionResponse{
			Message: "Transaction successfully registered",
			ID:      transaction.ID,
		})

	}
}

func isTransactionRegistered(ctx context.Context, id string) (bool, error) {
	user, err := repository.GetTradeById(ctx, id)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}

func decodeTransactionRequest(r *http.Request) (TransactionRequest, error) {
	var request TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func generateRandomID() (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
