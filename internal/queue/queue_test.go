package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue(10, "testQueue")
	if q == nil {
		t.Fatal("NewQueue did not create a new queue")
	}
}

func TestPushAndPop(t *testing.T) {
	q := NewQueue(2, "testQueue")
	msg1 := &Message{Value: map[string]interface{}{"message": "first"}}
	msg2 := &Message{Value: map[string]interface{}{"message": "second"}}

	q.Push(msg1)
	q.Push(msg2)

	poppedMsg1 := q.Pop()
	if !reflect.DeepEqual(poppedMsg1.Value, map[string]interface{}{"message": "first"}) {
		t.Errorf("Pop did not return the first pushed message")
	}

	poppedMsg2 := q.Pop()
	if !reflect.DeepEqual(poppedMsg2.Value, map[string]interface{}{"message": "second"}) {
		t.Errorf("Pop did not return the second pushed message")
	}
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue(10, "testQueue")
	if !q.IsEmpty() {
		t.Errorf("IsEmpty should return true for a new queue")
	}

	q.Push(&Message{Value: map[string]interface{}{"message": "test"}})
	if q.IsEmpty() {
		t.Errorf("IsEmpty should return false for a queue with messages")
	}
}

func TestInfo(t *testing.T) {
	q := NewQueue(10, "testQueue")
	q.Push(&Message{Value: map[string]interface{}{"message": "test"}})

	info := q.Info()
	if info.Name != "testQueue" || info.Presize != 10 || info.Size != 10 || info.Head != 0 || info.Tail != 1 || info.Count != 1 {
		t.Errorf("Info did not return the correct queue information")
	}
}
