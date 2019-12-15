package utils

import (
	"encoding/json"
	"net/http"

	"github.com/valentergs/books_backend/models"
)

//ResponseJSON will be exported ========================================
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//RespondWithError will be exported ====================================
func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
	return
}
