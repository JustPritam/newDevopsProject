package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var tasks []Task

func main() {
	http.HandleFunc("/tasks", handleTasks)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(tasks)
	case "POST":
		var task Task
		json.NewDecoder(r.Body).Decode(&task)
		task.ID = len(tasks) + 1
		tasks = append(tasks, task)
		json.NewEncoder(w).Encode(task)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
