package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"sample-ws/internal/database"
)

type QueryHandler struct {
	DB *sql.DB
}

func (h QueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllQueryData(h.DB)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
