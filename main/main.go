package main

import (
	"fmt"
	"pq/pq"
)

func main() {
	q := pq.NewPq(4)
	q.Insert(byte('e'), 1)
	q.Insert(byte('l'), 1)
	q.Insert(byte('o'), 3)
	q.Insert(byte('H'), 0)
	q.Insert(byte('l'), 2)
	var s string
	for !q.IsEmpty() {
		data, ok := q.Fetch().(byte)
		if ok {
			s = s + string(data)
		}
	}
	fmt.Println(s)
}
