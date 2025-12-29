package magiclink

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type MagicLinkRequest struct {
	Email string `json:"email"`
}

func handleMagicLinkRequest(w http.ResponseWriter, r *http.Request) {
	var req MagicLinkRequest

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))
	// log.Println("Received magic link request for email:", email)
	if email == "" || !strings.Contains(email, "@") {
		http.Error(w, `{"error": "Email not valid"}`, http.StatusBadRequest)
		return
	}

	// Check if email is authorized
	if err := checkEmailAuthorization(r.Context(), email); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s not authorized"}`, email), http.StatusUnauthorized)
		return
	}

	// send magic link email here

	if err := sendMagicLink(r.Context(), email); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// return success response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("✓ Magic link sent to %s. Check your inbox (link expires in %d minutes).", email, magicLinkExpiryMinutes),
	})



	// // generate magic link token
	// token, err := generateMagicLinkToken(email)
	// if err != nil {
	// 	http.Error(w, "Failed to generate magic link", http.StatusInternalServerError)
	// 	return
	// }
}




func RouterHandler(r chi.Router) {
	r.Post("/magic-link/request", handleMagicLinkRequest)
}