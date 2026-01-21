package main

import (
	"log"
	"net/http"

	"sample-ws/internal/config"
	"sample-ws/internal/database"
	"sample-ws/internal/handlers"
	"sample-ws/internal/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	secret := []byte(cfg.JWTSecret)

	// /auth
	http.Handle("/auth", handlers.AuthHandler{DB: db, Secret: secret})

	// /query (محمي بالـ middleware)
	query := handlers.QueryHandler{DB: db}
	http.Handle("/query", middleware.JWTAuth(secret, query))

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
