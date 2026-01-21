package database

import (
	"database/sql"

	"sample-ws/internal/models"
)

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	row := db.QueryRow(
		"SELECT id , username, password from users WHERE username = ?",
		username,
	)

	var u models.User
	// check if user exsit
	err := row.Scan(&u.ID, &u.User, &u.PasswordHash)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	//check if there an error
	if err != nil {
		return nil, err
	}
	// return user data
	return &u, nil
}
