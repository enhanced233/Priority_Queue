package pq

type QueuePriorN struct {
	fifo []*queueFIFO
	N    uint
}

func (q *QueuePriorN) Initialize(N uint) {
	q.N = N
	q.fifo = make([]*queueFIFO, N)
	for i := uint(0); i < q.N; i++ {
		q.fifo[i] = &queueFIFO{}
	}
}

func (q *QueuePriorN) IsEmpty() bool {
	for i := uint(0); i < q.N; i++ {
		if !q.fifo[i].isEmpty() {
			return false
		}
	}
	return true
}

func (q *QueuePriorN) Insert(data interface{}, priority uint) {
	if q.N == 0 {
		return
	}
	if priority > q.N {
		priority = q.N - 1
	}
	newNode := &node{data: data, priority: priority}
	q.fifo[priority].insert(newNode)
}

func (q *QueuePriorN) Pull() interface{} {
	for i := uint(0); i < q.N; i++ {
		if !q.fifo[i].isEmpty() {
			return q.fifo[i].pull()
		}
	}
	return nil
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
