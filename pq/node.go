package pq

type node struct {
	data     byte
	priority int
	next     *node
}

func (n *node) getData() byte {
	return n.data
}

func (n *node) push(other *node) *node {
	if other.priority < n.priority {
		other.next = n
		return other
	} else {
		if n.next == nil {
			n.next = other
		} else {
			n.next = n.next.push(other)
		}
		return n
	}
}
