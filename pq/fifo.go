package pq

type node struct {
	data     interface{}
	priority uint
	next     *node
}

type queueFIFO struct {
	head *node
	tail *node
}

func (q *queueFIFO) isEmpty() bool {
	return q.head == nil
}

func (q *queueFIFO) insert(newNode *node) {
	if q.isEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *queueFIFO) pull() interface{} {
	if q.isEmpty() {
		return nil
	}
	data := q.head.data
	q.head = q.head.next
	return data
}
