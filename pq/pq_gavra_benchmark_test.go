package pq

import (
	"math/rand"
	"testing"
)

func BenchmarkGPQ_Write(b *testing.B) {
	b.StopTimer()
	q := New(b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.Write(High, true)
	}
}

func BenchmarkGPQ_Read(b *testing.B) {
	b.StopTimer()
	q := New(b.N)
	for i := 0; i < b.N; i++ {
		q.Write(High, true)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = q.Read()
	}
}

func BenchmarkGPQ_WriteParallel(b *testing.B) {
	b.StopTimer()
	q := New(b.N)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Write(High, true)
		}
	})
}

func BenchmarkGPQ_ReadParallel(b *testing.B) {
	b.StopTimer()
	q := New(b.N)
	for i := 0; i < b.N; i++ {
		q.Write(High, true)
	}
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = q.Read()
		}
	})
}

func BenchmarkGPQ_WriteParallelRandom(b *testing.B) {
	b.StopTimer()
	const k = 3
	q := New(b.N)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Write(Priority(rand.Intn(k)), true)
		}
	})
}

func BenchmarkGPQ_ReadParallelRandom(b *testing.B) {
	b.StopTimer()
	const k = 3
	q := New(b.N)
	for i := 0; i < b.N; i++ {
		q.Write(Priority(rand.Intn(k)), true)
	}
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = q.Read()
		}
	})
}
