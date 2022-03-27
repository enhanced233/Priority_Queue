package pq

type QueuePriorN struct {
	fifo     map[uint]*queueFIFO
	keyOrder []uint
}

func (q *QueuePriorN) IsEmpty() bool {
	for key, value := range q.fifo {
		if q.keyExists(key) && !value.isEmpty() {
			return false
		}
	}
	return true
}

func (q *QueuePriorN) keyExists(priority uint) bool {
	_, ok := q.fifo[priority]
	return ok
}

func (q *QueuePriorN) Insert(data interface{}, priority uint) {
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
}

func (q *QueuePriorN) Pull() interface{} {
	for i := 0; i < len(q.keyOrder); i++ {
		priority := q.keyOrder[i]
		if q.keyExists(priority) && !q.fifo[priority].isEmpty() {
			return q.fifo[priority].pull()
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
