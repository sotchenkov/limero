package queue

type Message struct {
	Value string `json:"msg"`
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int, name string) *Queue {
	return &Queue{
		name: name,
		msgs: make([]*Message, size),
		size: size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	name  string
	msgs  []*Message
	size  int
	head  int
	tail  int
	count int
}

// Push adds a msg to the queue.
func (q *Queue) Push(n *Message) {
	if q.head == q.tail && q.count > 0 {
		msgs := make([]*Message, len(q.msgs)+q.size)
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

// func main() {
// 	q := NewQueue(1)
// 	q.Push(&Node{2})
// 	q.Push(&Node{4})
// 	q.Push(&Node{6})
// 	q.Push(&Node{8})
// 	fmt.Println(q.Pop(), q.Pop(), q.Pop(), q.Pop())
// }
