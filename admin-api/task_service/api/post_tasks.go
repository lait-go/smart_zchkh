package api

import (
	"admin-api/task_service/internal/auth"
	"admin-api/task_service/internal/httpx"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only POST allowed")
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid form", err.Error())
		return
	}

	task := auth.Task{
		AccountId:   r.FormValue("account_id"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Category:    r.FormValue("category"),
		Status:      "created",
		CreatedAt:   time.Now().Format(time.RFC3339),
		History: []auth.HistoryItem{
			{
				Timestamp: time.Now().Format(time.RFC3339),
				Action:    "created",
				User:      "frontend",
			},
		},
	}

	files := r.MultipartForm.File["attachments"]
	for _, fileHeader := range files {
		src, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer src.Close()

		os.MkdirAll("./uploads", os.ModePerm)
		dstPath := filepath.Join("./uploads", fileHeader.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			continue
		}
		defer dst.Close()
		io.Copy(dst, src)

		task.Attachments = append(task.Attachments, auth.Attachment{
			Type: "file",
			Url:  "/uploads/" + fileHeader.Filename,
		})
	}

	if err := auth.CreateTask(&task); err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to save task", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "task created",
		"id":      task.ID,
	})
}
