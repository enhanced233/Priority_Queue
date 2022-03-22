package pq

type node struct {
	data     interface{}
	priority int
}

type Queue struct {
	NodeList []*node
}

func (q *Queue) IsEmpty() bool {
	return len(q.NodeList) == 0
}

func searchPosition(list []*node, priority int) int {
	n := len(list)
	max := n
	min := 0
	current := n / 2
	for (max - min) > 0 {
		if list[current].priority <= priority {
			min = current
			current += (max - min) / 2
		} else {
			max = current
			current -= (max - min) / 2
		}
		if (max - min) == 1 {
			if list[min].priority <= priority {
				return max
			}
			return min
		}
	}
	return max
}

func (q *Queue) Insert(data interface{}, priority int) {
	newNode := &node{data: data, priority: priority}
	if q.IsEmpty() {
		q.NodeList = append(q.NodeList, newNode)
	} else {
		pos := searchPosition(q.NodeList, priority)
		q.NodeList = append(q.NodeList[:pos], append([]*node{newNode}, q.NodeList[pos:]...)...)
	}
}

func (q *Queue) Pull() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.NodeList[0].data
	q.NodeList = q.NodeList[1:]
	return data
}
