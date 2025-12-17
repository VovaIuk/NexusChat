package get_chathistory

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var usecase *Usecase

func HTTP_V1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	chatIDparam := r.PathValue("id")
	if chatIDparam == "" {
		http.Error(w, `{"error": "chat_id is required"}`, http.StatusBadRequest)
		return
	}

	chatID, err := strconv.Atoi(chatIDparam)
	if err != nil || chatID <= 0 {
		http.Error(w, `{"error": "invalid chat_id format"}`, http.StatusBadRequest)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, `{"error": "missing or invalid authorization header"}`, http.StatusUnauthorized)
		return
	}
	jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

	input := Input{
		ChatID:   chatID,
		JWTToken: jwtToken,
	}

	output, err := usecase.GetChatHistory(ctx, input)
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
