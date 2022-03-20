package pq

type Queue struct {
	head *node
}

func (q *Queue) Insert(data interface{}, priority int) {
	newNode := &node{data: data, priority: priority}
	if q.IsEmpty() {
		q.head = newNode
	} else {
		q.head = q.head.push(newNode)
	}
}

func (q *Queue) Pull() interface{} {
	if q.IsEmpty() {
		panic("Error: the queue is empty, cannot pull data!!")
	}
	data := q.head.getData()
	q.head = q.head.next
	return data
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}
