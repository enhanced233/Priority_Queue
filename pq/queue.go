package pq

import "sync"

type Queue struct {
	mutex    sync.Mutex
	fifo     map[uint]*queueFIFO
	keyOrder []uint
}

func (q *Queue) IsEmpty() bool {
	q.mutex.Lock()
	for key, value := range q.fifo {
		if q.keyExists(key) && !value.isEmpty() {
			q.mutex.Unlock()
			return false
		}
	}
	q.mutex.Unlock()
	return true
}

func (q *Queue) keyExists(priority uint) bool {
	_, ok := q.fifo[priority]
	return ok
}

func (q *Queue) Insert(data interface{}, priority uint) {
	q.mutex.Lock()
	if q.fifo == nil {
		q.fifo = make(map[uint]*queueFIFO)
		q.keyOrder = []uint{}
	}
	newNode := &node{data: data, priority: priority}
	if !q.keyExists(priority) {
		q.fifo[priority] = &queueFIFO{}
		pos := searchPosition(q.keyOrder, priority)
		q.keyOrder = append(q.keyOrder[:pos], append([]uint{priority}, q.keyOrder[pos:]...)...)
	}
	q.fifo[priority].insert(newNode)
	q.mutex.Unlock()
}

func (q *Queue) Pull() interface{} {
	q.mutex.Lock()
	for i := 0; i < len(q.keyOrder); i++ {
		priority := q.keyOrder[i]
		if q.keyExists(priority) && !q.fifo[priority].isEmpty() {
			q.mutex.Unlock()
			return q.fifo[priority].pull()
		}
	}
	q.mutex.Unlock()
	return nil
}

func searchPosition(list []uint, priority uint) int {
	n := len(list)
	max := n
	min := 0
	current := n / 2
	for (max - min) > 0 {
		if list[current] <= priority {
			min = current
			current += (max - min) / 2
		} else {
			max = current
			current -= (max - min) / 2
		}
		if (max - min) == 1 {
			if list[min] <= priority {
				return max
			}
			return min
		}
	}
	return max
}
