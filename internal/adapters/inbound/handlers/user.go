package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cristiangar0398/leal-rewards/internal/adapters/outbound/repository"
	"github.com/cristiangar0398/leal-rewards/internal/models"
	"github.com/cristiangar0398/leal-rewards/internal/server"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	request  SignUpRequest
	response SignUpResponse
)

type SignUpRequest struct {
	Cc       string `json:"document"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id string `json:"id"`
	Cc string `json:"document"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		request, err := decodeSignUpRequest(r)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		userID, err := generateUserID()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := createUser(r.Context(), request, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id: user.Id,
			Cc: user.Cc,
		})
	}
}

func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := repository.GetUserByDocument(r.Context(), request.Cc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			http.Error(w, "invalid credential ", http.StatusUnauthorized)
			return
		}

		userResponse, err := repository.GetUserDetail(r.Context(), request.Cc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var totalCashback float64
		for _, cashback := range userResponse.Cashback {
			totalCashback += cashback.Amount
		}

		userDetail := models.UserDetail{
			Id:            userResponse.Id,
			Cc:            userResponse.Cc,
			Points:        userResponse.Points,
			Cashback:      userResponse.Cashback,
			TotalCashback: totalCashback,
			//Trades:        userResponse.Trades,
		}

		w.Header().Set("content-type", "aaplication/json")
		json.NewEncoder(w).Encode(userDetail)
	}
}

func decodeSignUpRequest(r *http.Request) (SignUpRequest, error) {
	var request SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func generateUserID() (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func createUser(ctx context.Context, request SignUpRequest, userID string) (*models.User, error) {
	hashCostStr := os.Getenv("HASH_COST")
	hashCost, err := strconv.Atoi(hashCostStr)
	if err != nil {
		return nil, fmt.Errorf("Error al convertir HASH_COST a entero:", err)
	}

	//logica de creacion de usuario
	isRegistered, err := isDocumentRegistered(ctx, request.Cc)
	if err != nil {
		return nil, err
	}
	if isRegistered {
		return nil, fmt.Errorf("Document already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), hashCost)

	user := &models.User{
		Cc:       request.Cc,
		Password: string(hashedPassword),
		Id:       userID,
	}
	err = repository.InsertUser(ctx, user)
	return user, err
}

func isDocumentRegistered(ctx context.Context, document string) (bool, error) {
	user, err := repository.GetUserByDocument(ctx, document)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
