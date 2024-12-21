package main

import (
	"log"
	"net/http"

	"github.com/ZolotarevAlexandr/yl_sprint_1_final/calculator_server/internal/server"
)

func main() {
	mux := http.NewServeMux()
	calculateHandler := http.HandlerFunc(server.CalculateHandler)
	mux.Handle("/api/v1/calculate", server.ErrorHandlingMiddleware(server.LoggingMiddleware(calculateHandler)))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
