package auth

import (
	"admin-api/task_service/internal/logger"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var taskFile = "./data/tasks.json"
var console = &logger.Console{}

func LoadTasks() ([]Task, error) {
	file, err := os.Open(taskFile)
	if err != nil {
		console.Errorf("Failed to open file: %v", err)
		return nil, err
	}
	defer file.Close()
	console.Log("File opened successfully")

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		console.Errorf("Failed to decode JSON: %v", err)
		return nil, err
	}

	return tasks, nil
}

func SaveTask(tasks []Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		console.Errorf("Failed to create file: %v", err)
		return err
	}
	defer file.Close()
	console.Log("File created successfully")

	return json.NewEncoder(file).Encode(tasks)
}

func CreateTask(newTask *Task) error {
	tasks, err := LoadTasks()
	if err != nil {
		console.Errorf("Failed to load tasks: %v", err)
		return err
	}
	console.Log("Tasks loaded successfully")

	newTask.ID = fmt.Sprintf("TASK-%d", time.Now().UnixNano())
	tasks = append(tasks, *newTask)
	return SaveTask(tasks)
}
