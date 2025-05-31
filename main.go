package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var tasks = []Task{
	{ID: 1, Text: "First task"},
	{ID: 2, Text: "Second task"},
}

func main() {
	http.HandleFunc("/tasks", handleTasks)
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(tasks)
	case "POST":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task.ID = len(tasks) + 1
		tasks = append(tasks, task)
		json.NewEncoder(w).Encode(task)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
