package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "morse-app ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Starting server on :8080")
	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
