package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"sample-ws/internal/database"
	"sample-ws/internal/utils"
)

type AuthHandler struct {
	DB     *sql.DB
	Secret []byte
}

type authReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResp struct {
	Token string `json:"token"`
}

func (h AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req authReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	u, err := database.GetUserByUsername(h.DB, req.Username)
	if err != nil || u == nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPassword(u.PasswordHash, req.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(h.Secret, u.User)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authResp{Token: token})
}
