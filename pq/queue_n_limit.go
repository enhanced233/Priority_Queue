package pq

type QueueLimited struct {
	head *node
}

func (q *QueueLimited) IsEmpty() bool {
	return q.head == nil
}

func (q *QueueLimited) Insert(data interface{}, priority uint) {
	newNode := &node{data: data, priority: priority}
	if q.IsEmpty() {
		q.head = newNode
	}
}

func (q *QueueLimited) Pull() interface{} {

	return nil
}
