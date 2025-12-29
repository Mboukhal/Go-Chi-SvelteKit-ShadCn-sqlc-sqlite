package googleoauth

import (
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes sets up the OAuth routes on the given router.
func RegisterAuthRoutes(r chi.Router) {
	r.Route("/auth/google", func(r chi.Router) {
		r.Get("/login", handleOAuthLogin)
		r.Get("/callback", handleOAuthCallback)
	})
	r.Get("/auth/logout", LogoutHandler)
}
