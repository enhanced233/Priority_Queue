package main

import (
	"fmt"
	"pq/pq"
)

func main() {
	q := pq.Queue{}
	mapData := map[byte][]int{
		'H': {1},
		'e': {2},
		'l': {3, 3, 9},
		'o': {4, 7},
		' ': {5},
		'W': {6},
		'r': {8},
		'd': {10},
	}
	for i, v := range mapData {
		for j := 0; j < len(v); j++ {
			q.Insert(i, v[j])
		}

	}
	var s string
	for !q.IsEmpty() {
		s = s + string(q.Pull())
	}
	fmt.Println(s)
}
