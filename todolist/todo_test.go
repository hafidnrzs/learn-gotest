package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
