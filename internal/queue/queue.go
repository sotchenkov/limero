package queue

import (
	"sync"

	"github.com/sotchenkov/limero/internal/lib/response"
)

type Message struct {
	Value map[string]interface{} `json:"value"`
}

// NewQueue returns a new queue with the given initial presize.
func NewQueue(presize int, name string) *Queue {
	return &Queue{
		name:    name,
		msgs:    make([]*Message, presize),
		presize: presize,
	}
}

// Queue is a basic FIFO queue based on a circular list that represizes as needed.
type Queue struct {
	name    string
	msgs    []*Message
	presize int
	head    int // first element index
	tail    int // index where the next element will be placed when the Push is called
	count   int
	Mu      sync.Mutex
}

// Push adds a msg to the queue.
func (q *Queue) Push(n *Message) {
	if q.head == q.tail && q.count > 0 {
		msgs := make([]*Message, len(q.msgs)+q.presize)
		copy(msgs, q.msgs[q.head:])
		copy(msgs[len(q.msgs)-q.head:], q.msgs[:q.head])
		q.head = 0
		q.tail = len(q.msgs)
		q.msgs = msgs
	}
	q.msgs[q.tail] = n
	q.tail = (q.tail + 1) % len(q.msgs)
	q.count++
}

// Pop removes and returns a msg from the queue in first to last order.
func (q *Queue) Pop() *Message {
	if q.count == 0 {
		return nil
	}
	node := q.msgs[q.head]
	q.head = (q.head + 1) % len(q.msgs)
	q.count--
	return node
}

func (q *Queue) Info() *response.QueueInfo {
	return &response.QueueInfo{Name: q.name, Presize: q.presize, Size: len(q.msgs), Head: q.head, Tail: q.tail, Count: q.count}
}

func (q *Queue) IsEmpty() bool {
	return q.count == 0
}
