package login_user

import (
	"context"
	"encoding/json"
	"net/http"
)

var usecase *Usecase

func HTTP_V1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input: unable to parse JSON", http.StatusBadRequest)
		return
	}

	output, err := usecase.RegisterUser(context.Background(), input)
	if err != nil {
		http.Error(w, "Registration failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
