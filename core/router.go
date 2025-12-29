package router

import (
	oauth "github.com/Mboukhal/FactoryBase/core/auth/google"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes sets up the OAuth routes on the given router.
func RegisterRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		oauth.RegisterAuthRoutes(r)
		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(oauth.AuthMiddleware) // only for APIs
			// r.Get("/profile", oauth.HandleProfile)
			// TODO: check revoked tokens
		})
	})
}
