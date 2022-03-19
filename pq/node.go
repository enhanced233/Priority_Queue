package pq

type node struct {
	data     int8
	priority int
	next     *node
}

func (n *node) getPriority() int {
	return n.priority
}

func (n *node) getData() int8 {
	return n.data
}

func (n *node) push(other *node) *node {
	if other.priority < n.priority {
		other.next = n
		return other
	} else {
		n.next = n.next.push(other)
		return n
	}
}

func (n *node) pop() int8 {
	return n.data
}
