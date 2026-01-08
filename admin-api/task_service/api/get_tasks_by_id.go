package api

import (
	"admin-api/task_service/internal/auth"
	"admin-api/task_service/internal/httpx"
	"encoding/json"
	"net/http"
	"strings"
)

func GetTasksHandlerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		console.Errorf("Method %s not allowed", r.Method)
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET allowed")
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 || parts[4] == "" {
		console.Errorf("Invalid path: %s", r.URL.Path)
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid path", "Expected /api/tasks/{id}")
		return
	}
	taskID := parts[4]

	tasks, err := auth.LoadTasks()
	if err != nil {
		console.Errorf("Failed to load tasks: %v", err)
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load tasks", err.Error())
		return
	}
	
	for _, t := range tasks {
		if taskID == t.ID {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	
	httpx.NewJSONError(w, http.StatusNotFound, "Task not found", "No task with given ID")
}
