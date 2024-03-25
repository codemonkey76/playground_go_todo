package server

import (
	"log"
	"net/http"
	"todo/routes"
)

func StartServer() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
