package controllers

import (
	"encoding/json"
	u "go-petshop/internal/models/usuario"
	"go-petshop/internal/services"
	"net/http"
)

type UsuarioController struct {
	service services.UsuarioService
}

func NewUsuarioController(service services.UsuarioService) UsuarioController {
	return UsuarioController{service}
}

func (controller *UsuarioController) CriarUsuario(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var usuario u.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	usuarioCriado, err := controller.service.Create(ctx, &usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuarioCriado)
}
