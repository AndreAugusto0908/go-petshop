package router

import (
	"go-petshop/internal/controllers"

	"github.com/go-chi/chi/v5"
)

func setupUsuarioRoutes(r chi.Router, controller *controllers.UsuarioController) {
	r.Route("/v1/usuario", func(r chi.Router) {
		r.Post("/", controller.CriarUsuario)
	})
}
