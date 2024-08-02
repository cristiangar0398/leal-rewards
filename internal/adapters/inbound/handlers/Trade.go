package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cristiangar0398/leal-rewards/internal/adapters/outbound/repository"
	"github.com/cristiangar0398/leal-rewards/internal/models"
	"github.com/cristiangar0398/leal-rewards/internal/server"
	"github.com/segmentio/ksuid"
)

type TradeRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

type TradeResponse struct {
	Message   string `json:"message"`
	Id        string `json:"id"`
	TradeName string `json:"trade_name"`
	UserID    string `json:"user_id"`
}

func SignUpTradeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		request, err := decodeSignUpTradeRequest(r)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		tradeID, err := generateTradeID()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		trade, err := createTrade(r.Context(), request, request.UserId, tradeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-Type", "application/json")
		json.NewEncoder(w).Encode(TradeResponse{
			Message:   "Comercio Registrado con exito",
			Id:        trade.Id,
			TradeName: trade.TradeName,
			UserID:    trade.UserID,
		})
	}
}

func createTrade(ctx context.Context, request TradeRequest, userID string, tradeID string) (*models.Trade, error) {

	//logica de creacion de usuario
	isRegistered, err := isTradeRegistered(ctx, tradeID)
	if err != nil {
		return nil, err
	}
	if isRegistered {
		return nil, fmt.Errorf("Trade already registered")
	}

	trade := &models.Trade{
		Id:        tradeID,
		UserID:    userID,
		TradeName: request.Name,
	}
	err = repository.InsertTrade(ctx, trade)
	return trade, err
}

func isTradeRegistered(ctx context.Context, id string) (bool, error) {
	user, err := repository.GetTradeById(ctx, id)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}

func decodeSignUpTradeRequest(r *http.Request) (TradeRequest, error) {
	var request TradeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func generateTradeID() (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
