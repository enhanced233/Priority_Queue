package pq

import (
	"fmt"
)

func main() {
	q := Queue{}
	mapData := map[int8]int{
		'H': 1,
		'e': 2,
		'l': 2,
		'o': 3}
	for i, v := range mapData {
		q.Insert(i, v)
	}
	for !q.IsEmpty() {
		fmt.Println(q.Pull())
	}
}
