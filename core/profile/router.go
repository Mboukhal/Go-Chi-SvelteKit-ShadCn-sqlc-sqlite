package profile

import (
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes sets up the OAuth routes on the given router.
func RegisterAuthRoutes(r chi.Router) {
	r.Route("/profile", func(r chi.Router) {
		r.Get("/get-user/{userId}", getUserProfileHandler)
	})
}
