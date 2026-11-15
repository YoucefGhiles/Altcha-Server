package handlers

import (
	"encoding/json"
	"net/http"

	"altcha-server/service"
)

func ChallengeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "GET uniquement", http.StatusMethodNotAllowed)
		return
	}

	challenge, err := service.CreateChallenge()

	if err != nil {
		http.Error(
			w,
			"Erreur lors de la création du challenge",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(challenge)
}