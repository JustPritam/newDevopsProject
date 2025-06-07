package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

type Task struct {
	Title string
}

var (
	tasks []Task
	mutex sync.Mutex
)

var tmpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head><title>Simple To-Do App</title></head>
<body>
	<h1>To-Do List</h1>
	<form method="POST" action="/add">
		<input type="text" name="title" placeholder="New Task">
		<button type="submit">Add</button>
	</form>
	<ul>
		{{range .}}
			<li>{{.Title}}</li>
		{{end}}
	</ul>
</body>
</html>
`))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	tmpl.Execute(w, tasks)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		if title != "" {
			mutex.Lock()
			tasks = append(tasks, Task{Title: title})
			mutex.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)

	fmt.Println("To-Do app running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
