package main

import (
	"fmt"
	"pq/pq"
)

func main() {
	q := pq.NewQueuePriorN(10)
	mapData := map[byte][]int{
		'H': {0},
		'e': {1},
		'l': {2, 2, 8},
		'o': {3, 6},
		' ': {4},
		'W': {5},
		'r': {7},
		'd': {9},
	}
	for i, v := range mapData {
		for j := 0; j < len(v); j++ {
			q.Insert(i, uint(v[j]))
		}

	}
	var s string
	for !q.IsEmpty() {
		s = s + string(q.Fetch().(byte))
	}
	fmt.Println(s)
}
