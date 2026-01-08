package main

import (
	"admin-api/task_service/api"
	"admin-api/task_service/internal/httpx"
	"admin-api/task_service/internal/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			api.GetTasksHandler(w, r)
		case http.MethodPost:
			api.PostTasksHandler(w, r)
		default:
			httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET and POST allowed")
		}
	})
	mux.HandleFunc("/api/v1/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			api.GetTasksHandlerByID(w, r)
		case http.MethodPut:
			api.PutTasksById(w, r)
		default:
			httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET and PUT allowed")
		}
	})
	mux.HandleFunc("/api/v1/tasks/comment/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			api.GetTaskCommentHandler(w, r)
		default:
			httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET allowed")
		}
	})

	mux.HandleFunc("/api/v1/tasks/comment", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			api.PostTaskCommentHandler(w, r)
		default:
			httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only POST allowed")
		}
	})

	handler := middleware.CorsMiddleware(mux)
	http.ListenAndServe(":8081", handler)
}