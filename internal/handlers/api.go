package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
"go_tutorials/webserver/internal/middleware"
)


func Handler(r *chi.Mux) {
	// strip the slashes from the request (e.g. auth/api/ to auth/api)
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// middleware to check if the user is authenticated
		router.Use(middleware.Authorization)

		// get the balance of the user
		router.Get("/coins", GetCoinBalance)
	})
}