package main

import (
	"auth-api/api"
	"auth-api/internal/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/login", api.LoginHandler)
	mux.HandleFunc("/api/register", api.RegisterHandler)

	handler := middleware.CorsMiddleware(mux)

	http.ListenAndServe(":8082", handler)
}