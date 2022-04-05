package pq

import "sync"

type QueuePriorN struct {
	fifo           []*FIFO
	priorities     uint
	mutex          sync.Mutex
	firstAvailable uint
}

func NewQueuePriorN(priorities uint) *QueuePriorN {
	if priorities == 0 {
		return nil
	}
	q := &QueuePriorN{priorities: priorities}
	q.fifo = make([]*FIFO, priorities)
	for i := uint(0); i < q.priorities; i++ {
		q.fifo[i] = &FIFO{}
	}
	q.firstAvailable = priorities
	return q
}

func (q *QueuePriorN) IsEmpty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for i := uint(0); i < q.priorities; i++ {
		if !q.fifo[i].isEmpty() {
			return false
		}
	}
	return true
}

func (q *QueuePriorN) Insert(data interface{}, priority uint) {
	if priority >= q.priorities {
		priority = q.priorities - 1
	}
	newNode := &node{data: data, priority: priority}
	q.mutex.Lock()
	q.fifo[priority].push(newNode)
	if q.firstAvailable > priority {
		q.firstAvailable = priority
	}
	q.mutex.Unlock()
}

func (q *QueuePriorN) Fetch() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for q.firstAvailable < q.priorities {
		if !q.fifo[q.firstAvailable].isEmpty() {
			return q.fifo[q.firstAvailable].pop()
		} else {
			q.firstAvailable++
		}
	}
	return nil
}
