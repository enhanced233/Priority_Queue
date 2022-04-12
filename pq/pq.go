package pq

import "sync"

type PriorityQueue struct {
	fifo            []*FIFO
	numOfPriorities uint
	mutex           sync.Mutex
	firstAvailable  uint // Points to first available priority in the pq.
}

func NewPq(numOfPriorities uint) *PriorityQueue {
	if numOfPriorities == 0 {
		return nil
	}
	q := &PriorityQueue{numOfPriorities: numOfPriorities}
	q.fifo = make([]*FIFO, numOfPriorities)
	q.firstAvailable = numOfPriorities
	for i := uint(0); i < q.numOfPriorities; i++ {
		q.fifo[i] = &FIFO{}
	}
	return q
}

func (q *PriorityQueue) IsEmpty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for i := uint(0); i < q.numOfPriorities; i++ {
		if !q.fifo[i].isEmpty() {
			return false
		}
	}
	return true
}

// Insert - Inserts data into the desired priority queue.
func (q *PriorityQueue) Insert(data interface{}, priority uint) {
	if priority >= q.numOfPriorities {
		priority = q.numOfPriorities - 1
	}
	newNode := &node{data: data, priority: priority}
	q.mutex.Lock()
	q.fifo[priority].push(newNode)
	if q.firstAvailable > priority {
		q.firstAvailable = priority
	}
	q.mutex.Unlock()
}

// Fetch - Fetches the data from the first priority available .
func (q *PriorityQueue) Fetch() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for q.firstAvailable < q.numOfPriorities {
		if !q.fifo[q.firstAvailable].isEmpty() {
			return q.fifo[q.firstAvailable].pop()
		} else {
			q.firstAvailable++
		}
	}
	return nil
}
