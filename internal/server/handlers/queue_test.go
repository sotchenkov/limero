package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrongMenthod(t *testing.T) {
	// Test for Method Not Allowed
	req, err := http.NewRequest("GET", "/createQueue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	createQueue(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestCreateQueueWithoutName(t *testing.T) {
	req, err := http.NewRequest("PUT", "/createQueue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	createQueue(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestWrongQueueSizeType(t *testing.T) {
	req, err := http.NewRequest("PUT", "/createQueue?name=test&size=abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	createQueue(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestQueueAlreadyExists(t *testing.T) {
	// Assuming queues is a global variable or accessible in the test scope

	req, err := http.NewRequest("PUT", "/createQueue?name=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	createQueue(rr, req)

	rr = httptest.NewRecorder()
	createQueue(rr, req)

	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}
}

func TestSuccessfulQueueCreation(t *testing.T) {
	req, err := http.NewRequest("PUT", "/create-queue?name=newQueue&size=10", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	createQueue(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

}
