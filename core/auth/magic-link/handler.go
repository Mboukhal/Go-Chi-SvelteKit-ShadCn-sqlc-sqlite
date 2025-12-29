package magiclink

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)




func handleMagicLinkRequest(w http.ResponseWriter, r *http.Request) {
	// get email from form
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}



	// // generate magic link token
	// token, err := generateMagicLinkToken(email)
	// if err != nil {
	// 	http.Error(w, "Failed to generate magic link", http.StatusInternalServerError)
	// 	return
	// }
}




func RouterHandler(r chi.Router) {

	r.Post("/api/v1/auth/magic-link/request", handleMagicLinkRequest)
}