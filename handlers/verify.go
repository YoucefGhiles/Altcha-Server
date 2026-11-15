package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	altcha "github.com/altcha-org/altcha-lib-go/v2"

	"altcha-server/service"
)

type VerifyRequest struct {
	Payload string `json:"payload"`
}

type VerifyResponse struct {
	Verified bool   `json:"verified"`
	Reason   string `json:"reason,omitempty"`
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(VerifyResponse{
			Verified: false,
			Reason:   "method_not_allowed",
		})

		return
	}

	var request VerifyRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(VerifyResponse{
			Verified: false,
			Reason:   "invalid_json",
		})

		return
	}

	if request.Payload == "" {

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(VerifyResponse{
			Verified: false,
			Reason:   "missing_payload",
		})

		return
	}

	decoded, err := base64.StdEncoding.DecodeString(request.Payload)

	if err != nil {
		decoded, err = base64.RawStdEncoding.DecodeString(request.Payload)
	}

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(VerifyResponse{
			Verified: false,
			Reason:   "invalid_base64",
		})

		return
	}

	var payload altcha.Payload

	if err := json.Unmarshal(decoded, &payload); err != nil {

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(VerifyResponse{
			Verified: false,
			Reason:   "invalid_altcha_payload",
		})

		return
	}

	result, err := service.VerifyChallenge(
		payload.Challenge,
		payload.Solution,
	)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(VerifyResponse{
			Verified: false,
			Reason:   "verification_error",
		})

		return
	}

	if !result.Verified {

		reason := "verification_failed"

		if result.Expired {
			reason = "expired"
		}

		if result.InvalidSignature != nil && *result.InvalidSignature {
			reason = "invalid_signature"
		}

		if result.InvalidSolution != nil && *result.InvalidSolution {
			reason = "invalid_solution"
		}

		json.NewEncoder(w).Encode(VerifyResponse{
			Verified: false,
			Reason:   reason,
		})

		return
	}

	json.NewEncoder(w).Encode(VerifyResponse{
		Verified: true,
	})
}