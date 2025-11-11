package controllers

import (
	"encoding/json"
	"go-petshop/internal/dtos"
	"go-petshop/internal/services"
	"net/http"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) AuthController {
	return AuthController{service: service}
}

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var loginRequest dtos.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	token, err := a.service.Login(ctx, &loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}
