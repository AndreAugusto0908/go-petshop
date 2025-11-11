package router

import (
	"go-petshop/internal/controllers"

	"github.com/go-chi/chi/v5"
)

func setupAuthRoutes(r chi.Router, controller *controllers.AuthController) {
	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/login", controller.Login)
	})
}
