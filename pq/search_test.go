package pq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	s := "Hello World!"
	q := Queue{}
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

func TestQueueConcurrency(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	line := "World"
	s := ""
	mul := 2000
	split := 10
	out := ""
	outCh := make(chan string, split)

	for i := 0; i < mul; i++ {
		s = s + line
	}
	length := len(s)
	N := 3 // number of priorities
	q := NewQueuePriorN(uint(N))
	for i := 0; i < split; i++ {
		go func(i int) {
			for j := i * (length / split); j < (i+1)*(length/split) || (i == split-1 && j < length); j++ {
				priority := uint(rand.Intn(N))
				q.Insert(s[j], priority)
			}
		}(i)
	}
	for i := 0; i < split; i++ {
		go func(i int) {
			strOut := ""
			for len(strOut) < length/split {
				data, ok := q.Pull().(byte)
				if ok {
					strOut = strOut + string(data)
				}
			}
			outCh <- strOut
		}(i)
	}
	for i := 0; i < split; i++ {
		out = out + <-outCh
	}
	count := map[byte]int{
		'W': 0,
		'o': 0,
		'r': 0,
		'l': 0,
		'd': 0}
	for i := 0; i < len(out); i++ {
		count[out[i]]++
	}
	check := true
	for _, v := range count {
		if v != mul {
			check = false
		}
	}
	if check {
		fmt.Println("Correct!")
	} else {
		fmt.Println("ERROR: Incorrect!")
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
