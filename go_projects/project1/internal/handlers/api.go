package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes)
	// r.Use(chimiddle.Logger)
	// r.Use(chimiddle.Recoverer)
	// r.Use(middleware.CorsMiddleware)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}