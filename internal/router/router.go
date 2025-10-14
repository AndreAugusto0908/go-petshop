package router

import (
	"go-petshop/internal/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	UsuarioController *controllers.UsuarioController
}

func NewRouter(usuarioController *controllers.UsuarioController) *Router {
	return &Router{
		UsuarioController: usuarioController,
	}
}

func (router *Router) Setup() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(r chi.Router) {
		setupUsuarioRoutes(r, router.UsuarioController)
	})
	return r
}
