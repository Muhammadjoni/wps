package routes

import (
	"backend/internal/controller"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
)

func ProductRoutes(r chi.Router, h *controller.ProductHandler, tokenAuth *jwtauth.JWTAuth) {
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(Authenticator)

		r.Post("/", h.CreateProduct)
		r.Get("/", h.ListProducts)
		r.Get("/{id}", h.GetProduct)
	})
}
