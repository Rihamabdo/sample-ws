package handlers

import (
	"database/sql"
	"net/http"

	"sample-ws/internal/database"
	"sample-ws/internal/utils"
)

type CreateUserHandler struct {
	DB *sql.DB
}

func (h CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := "admin2"
	password := "123456"

	hash, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "hash error", 500)
		return
	}

	if err := database.CreateUser(h.DB, username, hash); err != nil {
		http.Error(w, "db error", 500)
		return
	}

	w.Write([]byte("user created"))
}
