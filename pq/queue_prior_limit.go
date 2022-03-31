package pq

import "sync"

type QueuePriorN struct {
	fifo  []*queueFIFO
	N     uint
	mutex sync.Mutex
}

func NewQueuePriorN(N uint) *QueuePriorN {
	if N == 0 {
		return nil
	}
	q := &QueuePriorN{N: N}
	q.fifo = make([]*queueFIFO, N)
	for i := uint(0); i < q.N; i++ {
		q.fifo[i] = &queueFIFO{}
	}
	return q
}

func (q *QueuePriorN) IsEmpty() bool {
	//q.mutex.Lock()
	//defer q.mutex.Unlock()
	for i := uint(0); i < q.N; i++ {
		if !q.fifo[i].isEmpty() {
			return false
		}
	}
	return true
}

func (q *QueuePriorN) Insert(data interface{}, priority uint) {
	if priority >= q.N {
		priority = q.N - 1
	}
	newNode := &node{data: data, priority: priority}
	//q.mutex.Lock()
	q.fifo[priority].insert(newNode)
	//q.mutex.Unlock()
}

func (q *QueuePriorN) Pull() interface{} {
	//q.mutex.Lock()
	//defer q.mutex.Unlock()
	for i := uint(0); i < q.N; i++ {
		if !q.fifo[i].isEmpty() {
			return q.fifo[i].pull()
		}
	}
	return nil
}
