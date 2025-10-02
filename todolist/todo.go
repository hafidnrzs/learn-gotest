package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// The Task struct model
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// In-memory store for tasks
var (
	tasks = make(map[string]Task)
	mu    = sync.Mutex{} // handle concurrent process safely
)

// handler for POST /tasks
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Decode incoming JSON payload from the request body
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Generate a new unique ID for the task
	newTask.ID = uuid.New().String()
	newTask.Completed = false // new tasks are not completed by default

	// 3. Lock the mutex to safely write to the tasks map, then unlock it
	mu.Lock()
	tasks[newTask.ID] = newTask
	mu.Unlock()

	// 4. Set the "Content-Type" header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// 5. Set the HTTP status code to 201 Created
	w.WriteHeader(http.StatusCreated)

	// 6. Encode the newly created task as JSON and write it to the response
	json.NewEncoder(w).Encode(newTask)
}

func getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex to safely read from the tasks map
	mu.Lock()
	defer mu.Unlock()

	taskSlice := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		taskSlice = append(taskSlice, task)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(taskSlice)
}

func getTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := chi.URLParam(r, "id")
	task, ok := tasks[id]
	if !ok {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)
}

// main function to set up the router and start the server
func main() {
	r := chi.NewRouter()

	r.Post("/tasks", createTaskHandler)
	r.Get("/tasks", getAllTasksHandler)
	r.Get("/tasks/{id}", getTaskByIdHandler)

	log.Printf("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
