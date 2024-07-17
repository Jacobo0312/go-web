package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	appErrors "github.com/Jacobo0312/go-web/pkg/errors"
)

func ReadIdParam(r *http.Request) (int64, error) {
	param := r.PathValue("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
}

func RespondWithError(w http.ResponseWriter, err *appErrors.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": err.Message}); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
