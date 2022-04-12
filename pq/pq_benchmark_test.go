package pq

import (
	"math/rand"
	"testing"
)

const k = 3

func Benchmark_Write(b *testing.B) {
	b.StopTimer()
	q := NewPq(k)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.Insert(true, 0)
	}
}

func Benchmark_Read(b *testing.B) {
	b.StopTimer()
	q := NewPq(k)
	for i := 0; i < b.N; i++ {
		q.Insert(true, 0)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = q.Fetch()
	}
}

func Benchmark_WriteParallel(b *testing.B) {
	b.StopTimer()
	q := NewPq(k)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Insert(true, 0)
		}
	})
}

func Benchmark_ReadParallel(b *testing.B) {
	b.StopTimer()
	q := NewPq(k)
	for i := 0; i < b.N; i++ {
		q.Insert(true, 0)
	}
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = q.Fetch()
		}
	})
}

func Benchmark_WriteParallelRandom(b *testing.B) {
	b.StopTimer()
	q := NewPq(k)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Insert(true, uint(rand.Intn(k)))
		}
	})
}

func Benchmark_ReadParallelRandom(b *testing.B) {
	b.StopTimer()
	q := NewPq(k)
	for i := 0; i < b.N; i++ {
		q.Insert(true, uint(rand.Intn(k)))
	}
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = q.Fetch()
		}
	})
}
