package auth

import (
	"encoding/json"
	"os"
	"sync"
)

type Comment struct {
	TaskID   string `json:"task_id"`
	User     string `json:"user"`
	Message  string `json:"text"`
	Datetime string `json:"datetime"`
}

var commentFile = "./comments/comments.json"
var mu sync.Mutex

func LoadComment(taskID string) ([]Comment, error) {
	mu.Lock()
	defer mu.Unlock()

	var all []Comment

	file, err := os.ReadFile(commentFile)
	if err == nil {
		json.Unmarshal(file, &all)
	}

	var result []Comment
	for _, c := range all {
		if c.TaskID == taskID {
			result = append(result, c)
		}
	}
	return result, nil
}

func SaveComment(newComment Comment) error {
	mu.Lock()
	defer mu.Unlock()

	var all []Comment
	file, err := os.ReadFile(commentFile)
	if err == nil {
		json.Unmarshal(file, &all)
	}

	all = append(all, newComment)

	data, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(commentFile, data, 0644)
}
