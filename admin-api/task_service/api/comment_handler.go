package api

import (
	"admin-api/task_service/internal/auth"
	"encoding/json"
	"net/http"
	"strings"
)

func GetTaskCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/tasks/comment/")
	comments, err := auth.LoadComment(id)
	if err != nil {
		http.Error(w, "Комментарий не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func PostTaskCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var comment auth.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	err := auth.SaveComment(comment)
	if err != nil {
		http.Error(w, "Ошибка при сохранении", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
