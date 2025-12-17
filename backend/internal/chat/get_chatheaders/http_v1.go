package get_chatheaders

import (
	"encoding/json"
	"net/http"
	"strings"
)

var usecase *Usecase

func HTTP_V1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, `{"error": "missing or invalid authorization header"}`, http.StatusUnauthorized)
		return
	}
	jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

	input := Input{
		JWTToken: jwtToken,
	}

	output, err := usecase.GetChatHeaders(ctx, input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
	}
}
