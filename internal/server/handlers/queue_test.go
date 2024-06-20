package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateQueueWithoutName(t *testing.T) {
	req, err := http.NewRequest("PUT", "/queue", nil)
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
	req, err := http.NewRequest("PUT", "/queue?name=test&size=abc", nil)
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

	req, err := http.NewRequest("PUT", "/queue?name=test", nil)
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
	req, err := http.NewRequest("PUT", "/queue?name=newQueue&size=10", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	createQueue(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

}

func TestDeleteQueueWithoutName(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/queue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	deleteQueue(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestDeleteQueueNotFound(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/queue?name=nonexistentQueue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	deleteQueue(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestSuccessfulDeleteQueue(t *testing.T) {
	req, err := http.NewRequest("PUT", "/queue?name=newQueue&size=10", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	createQueue(rr, req)

	req, err = http.NewRequest("DELETE", "/queue?name=newQueue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	deleteQueue(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestQueueInfoWithoutName(t *testing.T) {
	req, err := http.NewRequest("GET", "/queue/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	queueInfo(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestQueueInfoNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/queue/nonexistentQueue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	queueInfo(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestSuccessfulQueueInfo(t *testing.T) {
	req, err := http.NewRequest("PUT", "/queue?name=testQueue&size=10", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	createQueue(rr, req)

	req, err = http.NewRequest("GET", "/queue/testQueue", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	queueInfo(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestQueuesList(t *testing.T) {
	req, err := http.NewRequest("GET", "/queue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Queues(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
