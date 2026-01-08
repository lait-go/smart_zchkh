package api

import (
	"admin-api/task_service/internal/auth"
	"admin-api/task_service/internal/httpx"
	"encoding/json"
	"net/http"
	"strings"
)

type TaskUpdatePayload struct {
	Status     string `json:"status"`
	ExecutorID string `json:"executor_id"`
}

func PutTasksById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/tasks/")

	var payload TaskUpdatePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Failed to decode JSON", err.Error())
		return
	}

	tasks, err := auth.LoadTasks()
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load tasks", err.Error())
		return
	}

	updated := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = payload.Status
			tasks[i].ExecutorID = payload.ExecutorID
			updated = true
			break
		}
	}

	if !updated {
		httpx.NewJSONError(w, http.StatusNotFound, "Task not found", "Task with ID "+id+" not found")
		return
	}

	if err := auth.SaveTask(tasks); err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to save tasks", err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}
