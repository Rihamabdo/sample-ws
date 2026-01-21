package database

import (
	"database/sql"

	"sample-ws/internal/models"
)

func GetAllQueryData(db *sql.DB) ([]models.QueryData, error) {
	rows, err := db.Query("SELECT id, title, body FROM query_data ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.QueryData
	for rows.Next() {
		var d models.QueryData
		if err := rows.Scan(&d.ID, &d.Title, &d.Body); err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, rows.Err()

}
