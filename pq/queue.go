package pq

type Queue struct {
	head *node
}

func (q *Queue) Insert(data int8, priority int) {
	newNode := &node{data: data, priority: priority}
	if q.IsEmpty() {
		q.head = newNode
	} else {
		q.head.push(newNode)
	}
}

func (q *Queue) Pull() int8 {
	if q.IsEmpty() {
		panic("Error: the queue is empty, cannot pull data!!")
	}
	data := q.head.pop()
	q.head = q.head.next
	return data
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}
