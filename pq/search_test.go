package pq

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	N          = 3    // # of priorities
	selectFunc = "RW" // "MW"/"MR"/"RW"
	dataSize   = 100000
	bins       = 10 // # of goroutines in parallel for MW/MR
)

func TestQueuePrior(t *testing.T) {
	n := N
	for ; n <= 81; n *= N {
		fmt.Println("N =", n)
		testQueueOrder(n, t)
		testQueueConcurrency(n, t)
	}
}

func testQueueOrder(n int, t *testing.T) {
	q := NewQueuePriorN(uint(n))
	num, count := 0, 0
	for i := 0; i < dataSize; i++ {
		priority := uint(i * n / dataSize)
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

func testQueueConcurrency(n int, t *testing.T) {
	q := NewQueuePriorN(uint(n))
	for i := 0; i < bins; i++ {
		go func() {
			in := false
			for j := 0; j < dataSize/bins; j++ {
				priority := uint(rand.Intn(n))
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

func benchmarkMW(n int, b *testing.B) {
	q := NewQueuePriorN(uint(n))
	end := make(chan int)
	b.StartTimer()
	for i := 0; i < bins; i++ {
		go func(end chan int) {
			in := false
			for j := 0; j < dataSize/bins; j++ {
				priority := uint(rand.Intn(n))
				q.Insert(in, priority)
				in = !in
			}
			end <- 0
		}(end)
	}
	for i := 0; i < bins; i++ {
		<-end
	}
}

func benchmarkMR(n int, b *testing.B) {
	q := NewQueuePriorN(uint(n))
	end := make(chan int)
	for i := 0; i < bins; i++ {
		go func(end chan int) {
			in := false
			for j := 0; j < dataSize/bins; j++ {
				priority := uint(rand.Intn(n))
				q.Insert(in, priority)
				in = !in
			}
			end <- 0
		}(end)
	}
	for i := 0; i < bins; i++ {
		<-end
	}
	b.StartTimer()
	for i := 0; i < bins; i++ {
		go func(end chan int) {
			count := 0
			for count < dataSize/bins {
				_, ok := q.Fetch().(bool)
				if ok {
					count++
				}
			}
			end <- 0
		}(end)
	}
	for i := 0; i < bins; i++ {
		<-end
	}
}

func benchmarkRW(n int, b *testing.B) {
	q := NewQueuePriorN(uint(n))
	b.StartTimer()
	in := false
	for i := 0; i < dataSize; i++ {
		priority := uint(rand.Intn(n))
		q.Insert(in, priority)
		in = !in
	}
	count := 0
	for count < dataSize {
		_, ok := q.Fetch().(bool)
		if ok {
			count++
		}
	}
}

func benchSelector(n int, b *testing.B) {
	b.StopTimer()
	switch selectFunc {
	case "RW":
		benchmarkRW(n, b)
	case "MW":
		benchmarkMW(n, b)
	case "MR":
		benchmarkMR(n, b)
	}
}

func BenchmarkQueueN(b *testing.B) {
	n := N // # of priorities
	benchSelector(n, b)
}

func BenchmarkQueue4N(b *testing.B) {
	n := 4 * N // # of priorities
	benchSelector(n, b)
}

func BenchmarkQueue40N(b *testing.B) {
	n := 40 * N // # of priorities
	benchSelector(n, b)
}
