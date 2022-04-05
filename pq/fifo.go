package pq

type node struct {
	data     interface{}
	priority uint
	next     *node
}

type FIFO struct {
	head *node
	tail *node
}

func (q *FIFO) isEmpty() bool {
	return q.head == nil
}

func (q *FIFO) push(newNode *node) {
	if q.isEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *FIFO) pop() interface{} {
	if q.isEmpty() {
		return nil
	}
	data := q.head.data
	q.head = q.head.next
	return data
}
