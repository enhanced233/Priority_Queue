package main

import (
	"fmt"
	"pq/pq"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	s := "Hello World!"
	q := pq.QueuePriorN{}
	N := 10 // number of priorities
	out := ""
	end := make(chan int)
	go func() {
		for i := 0; i < len(s); i++ {
			var num int
			if len(s) < N {
				num = 1
			} else {
				num = len(s) / N
			}
			priority := uint(i / num)
			q.Insert(s[i], priority)

		}
	}()
	go func() {
		for q.IsEmpty() {
			time.Sleep(1 * time.Microsecond)
		}
		for !q.IsEmpty() {
			out = out + string(q.Pull().(byte))
		}
		end <- 0
	}()
	<-end
	if s == out {
		fmt.Println("Correct!")
	} else {
		fmt.Println("ERROR: Incorrect!")
	}
}

func benchmarkQueue(N int, b *testing.B) {
	line := "Hello World!"
	s := ""
	for i := 0; i < 100; i++ {
		s = s + line
	}
	q := pq.QueuePriorN{}
	for n := 0; n < b.N; n++ {
		out := ""
		for i := 0; i < len(s); i++ {
			var num int
			if len(s) < N {
				num = 1
			} else {
				num = len(s) / N
			}
			priority := uint(i / num)
			q.Insert(s[i], priority)
		}
		for !q.IsEmpty() && line != out {
			out = out + string(q.Pull().(byte))
		}
		if line != out {

			panic(s + " - Incorrect output - " + out)
		}
	}
}

func BenchmarkQueue1(b *testing.B) {
	benchmarkQueue(1, b)
}

func BenchmarkQueue2(b *testing.B) {
	benchmarkQueue(2, b)
}

func BenchmarkQueue5(b *testing.B) {
	benchmarkQueue(5, b)
}

func BenchmarkQueue10(b *testing.B) {
	benchmarkQueue(10, b)
}

func BenchmarkQueue20(b *testing.B) {
	benchmarkQueue(20, b)
}

func BenchmarkQueue50(b *testing.B) {
	benchmarkQueue(50, b)
}

func BenchmarkQueue100(b *testing.B) {
	benchmarkQueue(100, b)
}
