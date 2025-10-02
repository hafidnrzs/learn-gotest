package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestCreateTaskHandler(t *testing.T) {
	// 1. Define the Task JSON payload sent in the request
	taskPayload := []byte(`{"title": "Learn TDD", "description": "Write a failing test first"}`)

	// 2. Create a new HTTP request with this payload
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskPayload))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// 3. Create a "ResponseRecorder" to record the response from the handler
	rr := httptest.NewRecorder()

	// 4. Create an HTTP handler from handler function. It will fail until the handler function are created
	handler := http.HandlerFunc(createTaskHandler)

	// 5. Serve the HTTP request to the handler
	handler.ServeHTTP(rr, req)

	// 6. Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 7. Decode the response body into Task struct to check its content
	var createdTask Task
	if err := json.NewDecoder(rr.Body).Decode(&createdTask); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	// 8. Check if the returned task has the correct title and non-empty ID
	if createdTask.Title != "Learn TDD" {
		t.Errorf("handler returned unexpected title: got %v want %v", createdTask.Title, "Learn TDD")
	}

	if createdTask.ID == "" {
		t.Errorf("handler returned empty ID")
	}

	if createdTask.Completed != false {
		t.Errorf("new task should not be completed")
	}
}

func TestGetAllTasksHandler(t *testing.T) {
	// 1. Prepare the in-memory store with some tasks
	mu.Lock()
	tasks = make(map[string]Task)
	mu.Unlock()

	tasks = map[string]Task{
		"0": {ID: "0", Title: "Get all tasks", Description: "Create test case for get all tasks", Completed: false},
		"1": {ID: "1", Title: "Learn TDD", Description: "Write a failing test first", Completed: false},
	}

	// 2. Create a new HTTP request to get all tasks
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// 3. Record the response
	rr := httptest.NewRecorder()

	// 4. Create the handler and execute it
	handler := http.HandlerFunc(getAllTasksHandler)
	handler.ServeHTTP(rr, req)

	// 5. Assert the outcome: status code, response body
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var allTasks []Task
	if err := json.NewDecoder(rr.Body).Decode(&allTasks); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if len(allTasks) != 2 {
		t.Errorf("handler returned wrong number of tasks: got %d want %d", len(allTasks), 2)
	}

	if allTasks[0].ID != "0" {
		t.Errorf("handler returned expected first task ID: got '%s' want '%s'", allTasks[0].ID, "0")
	}

	if allTasks[0].Title != "Get all tasks" {
		t.Errorf("handler returned expected first task Title: got '%s' want '%s'", allTasks[0].Title, "Get all tasks")
	}

	// Clear the tasks map
	mu.Lock()
	tasks = make(map[string]Task)
	mu.Unlock()
}

func TestGetTaskByIdHandler(t *testing.T) {
	// arrange
	mu.Lock()
	tasks = make(map[string]Task)
	mu.Unlock()

	tasks = map[string]Task{
		"0": {ID: "0", Title: "Get all tasks", Description: "Create test case for get all tasks", Completed: false},
		"1": {ID: "1", Title: "Learn TDD", Description: "Write a failing test first", Completed: false},
	}

	// action
	r := chi.NewRouter()
	r.Get("/tasks/{id}", getTaskByIdHandler)

	req, err := http.NewRequest("GET", "/tasks/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var task Task
	if err := json.NewDecoder(rr.Body).Decode(&task); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if task.ID != "1" {
		t.Errorf("handler returned unexpected ID: got %v want %v", task.ID, 1)
	}

	if task.Title != "Learn TDD" {
		t.Errorf("handler returned unexpected title: got %v want %v", task.Title, "Learn TDD")
	}

	if task.Completed != false {
		t.Errorf("handler returned unexpected completed status: got %t want %t", task.Completed, false)
	}
}
