package main

import "fmt"

func f() string {
	q := Queue{}
	mapData := map[byte]int{
		'H': 1,
		'e': 2,
		'l': 3,
		'o': 4}
	for i, v := range mapData {
		//if i == 'H' {
		//	fmt.Println("True", i, v)
		//}
		q.Insert(i, v)
	}
	var s string
	for !q.IsEmpty() {
		s = s + string(q.Pull())
	}
	return s
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
