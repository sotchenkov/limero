package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sotchenkov/limero/internal/queue"
)

func TestPushMsgWithoutQueueName(t *testing.T) {
	body := strings.NewReader(`{"value":{"message":"test message"}}`)
	req, err := http.NewRequest("POST", "/msg", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	PushMsg(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestPushMsgToNonExistentQueue(t *testing.T) {
	body := strings.NewReader(`{"value":{"message":"test message"}}`)
	req, err := http.NewRequest("POST", "/msg?qname=nonexistent", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	PushMsg(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestPushMsgWithUnsupportedType(t *testing.T) {
	req, err := http.NewRequest("PUT", "/queue?name=testQueue&size=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	createQueue(rr, req)

	body := strings.NewReader("not a json")
	req, err = http.NewRequest("POST", "/msg?qname=testQueue", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/plain")
	rr = httptest.NewRecorder()
	PushMsg(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestPushMsgWithEmptyBody(t *testing.T) {
	req, err := http.NewRequest("PUT", "/queue?name=testQueue&size=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	createQueue(rr, req)

	body := strings.NewReader(`{}`)
	req, err = http.NewRequest("POST", "/msg?qname=testQueue", body)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	PushMsg(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestSuccessfulMsgPush(t *testing.T) {
	req, err := http.NewRequest("PUT", "/queue?name=testQueue&size=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	createQueue(rr, req)

	body := strings.NewReader(`{"value": {"message": "test"}}`)
	req, err = http.NewRequest("POST", "/msg?qname=testQueue", body)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	PushMsg(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func BenchmarkPushMsgQueueSize1(b *testing.B) {
	// Setup
	queueName := "testQueueSize1"
	queues[queueName] = queue.NewQueue(1, queueName)
	defer delete(queues, queueName)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		body := strings.NewReader(`{"value":{"message":"message"}}`)
		req, _ := http.NewRequest("POST", "/msg?qname="+queueName, body)
		rr := httptest.NewRecorder()
		PushMsg(rr, req)
	}
}

func BenchmarkPushMsgQueueSize100(b *testing.B) {
	// Setup
	queueName := "testQueueSize100"
	queues[queueName] = queue.NewQueue(100, queueName)
	defer delete(queues, queueName)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		body := strings.NewReader(`{"value":{"message":"message"}}`)
		req, _ := http.NewRequest("POST", "/msg?qname="+queueName, body)
		rr := httptest.NewRecorder()
		PushMsg(rr, req)
	}
}

func BenchmarkPushMsgQueueSize100000(b *testing.B) {
	// Setup
	queueName := "testQueueSize100000"
	queues[queueName] = queue.NewQueue(100000, queueName)
	defer delete(queues, queueName)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		body := strings.NewReader(`{"value":{"message":"message"}}`)
		req, _ := http.NewRequest("POST", "/msg?qname="+queueName, body)
		rr := httptest.NewRecorder()
		PushMsg(rr, req)
	}
}

func TestPopMsgWithoutQueueName(t *testing.T) {
	req, err := http.NewRequest("GET", "/msg", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	PopMsg(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestPopMsgFromNonExistentQueue(t *testing.T) {
	req, err := http.NewRequest("GET", "/msg?qname=nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	PopMsg(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestPopMsgFromEmptyQueue(t *testing.T) {
	req, err := http.NewRequest("PUT", "/queue?name=emptyQueue&size=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	createQueue(rr, req)

	req, err = http.NewRequest("GET", "/msg?qname=emptyQueue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	PopMsg(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestSuccessfulMsgPop(t *testing.T) {
	// Setup: create a queue and push a message into it
	req, err := http.NewRequest("PUT", "/queue?name=testQueue&size=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	createQueue(rr, req)

	body := strings.NewReader(`{"value": {"message": "test"}}`)
	req, err = http.NewRequest("POST", "/msg?qname=testQueue", body)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	PushMsg(rr, req)

	// Test: pop the message
	req, err = http.NewRequest("GET", "/msg?qname=testQueue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	PopMsg(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v. Info: %s", status, http.StatusOK, rr.Body)
	}
}

func BenchmarkPopMsgQueueSize1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		queueName := "testQueueSize1"
		localQueue := queue.NewQueue(1, queueName)
		localQueue.Push(&queue.Message{Value: map[string]interface{}{"message": "message"}})

		for pb.Next() {
			req, err := http.NewRequest("GET", "/msg?qname="+queueName, nil)
			if err != nil {
				b.Error(err)
				return
			}
			rr := httptest.NewRecorder()
			PopMsg(rr, req)
		}
	})
}

func BenchmarkPopMsgQueueSize100(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		queueName := "testQueueSize100"
		localQueue := queue.NewQueue(100, queueName)
		localQueue.Push(&queue.Message{Value: map[string]interface{}{"message": "message"}})

		for pb.Next() {
			req, err := http.NewRequest("GET", "/msg?qname="+queueName, nil)
			if err != nil {
				b.Error(err)
				return
			}
			rr := httptest.NewRecorder()
			PopMsg(rr, req)
		}
	})
}

func BenchmarkPopMsgQueueSize100000(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		queueName := "testQueueSize100000"
		localQueue := queue.NewQueue(100000, queueName)
		localQueue.Push(&queue.Message{Value: map[string]interface{}{"message": "message"}})

		for pb.Next() {
			req, err := http.NewRequest("GET", "/msg?qname="+queueName, nil)
			if err != nil {
				b.Error(err)
				return
			}
			rr := httptest.NewRecorder()
			PopMsg(rr, req)
		}
	})
}
