package pq

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQueuePrior(t *testing.T) {
	dataSize := 10000
	bins := 10 // # of goroutines in parallel for MW/MR
	N := 3     // # of priorities
	for ; N <= 81; N *= 3 {
		fmt.Println("N =", N)
		testQueueOrder(dataSize, N, t)
		testQueueConcurrency(dataSize, bins, N, t)
	}
}

func testQueueOrder(dataSize, N int, t *testing.T) {
	q := NewQueuePriorN(uint(N))
	num, count := 0, 0
	for i := 0; i < dataSize; i++ {
		priority := uint(i * N / dataSize)
		q.Insert(num, priority)
		num++
	}
	for count < dataSize {
		out, ok := q.Fetch().(int)
		if !ok {
			continue
		}
		if out == count {
			count++
		} else {
			t.Errorf("Functional test - FAIL!")
			return
		}
	}
	fmt.Println("Functional test - PASS!")
}

func testQueueConcurrency(dataSize, bins, N int, t *testing.T) {
	q := NewQueuePriorN(uint(N))
	for i := 0; i < bins; i++ {
		go func() {
			in := false
			for j := 0; j < dataSize/bins; j++ {
				priority := uint(rand.Intn(N))
				q.Insert(in, priority)
				in = !in
			}
		}()
	}
	outTrue, outFalse := make(chan int), make(chan int)
	for i := 0; i < bins; i++ {
		go func() {
			count, trueCount, falseCount := 0, 0, 0
			for count < dataSize/bins {
				data, ok := q.Fetch().(bool)
				if ok {
					count++
					if data {
						trueCount++
					} else {
						falseCount++
					}
				}
			}
			outTrue <- trueCount
			outFalse <- falseCount
		}()
	}
	countTrue, countFalse := 0, 0
	for i := 0; i < bins; i++ {
		countTrue = countTrue + <-outTrue
		countFalse = countFalse + <-outFalse
	}
	if countTrue == countFalse && countTrue+countFalse == dataSize {
		fmt.Println("Concurrency test - PASS!")
	} else {
		t.Errorf("Concurrency test - FAIL!")
	}
}

func benchmarkQueue(N int, b *testing.B) {
	line := "Hello World!"
	s := ""
	b.StopTimer()
	for i := 0; i < 100; i++ {
		s = s + line
	}
	b.StartTimer()
	q := Queue{}
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
			out = out + string(q.Fetch().(byte))
		}
		if line != out {

			panic(s + " - Incorrect output - " + out)
		}
	}
}

func BenchmarkQueue1(b *testing.B) {
	benchmarkQueue(1, b)
}
