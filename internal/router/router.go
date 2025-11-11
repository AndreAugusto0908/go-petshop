package router

import (
	"go-petshop/internal/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	UsuarioController *controllers.UsuarioController
	AuthController    *controllers.AuthController
}

func NewRouter(usuarioController *controllers.UsuarioController, authControler *controllers.AuthController) *Router {
	return &Router{
		UsuarioController: usuarioController,
		AuthController:    authControler,
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
		setupAuthRoutes(r, router.AuthController)
	})
	return r
}
