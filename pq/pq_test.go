package pq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPQ_ZeroPriorities(t *testing.T) {
	q := NewQueuePriorN(uint(0))
	assert.Nil(t, q)
}

func TestPQ_Insert(t *testing.T) {
	q := NewQueuePriorN(uint(3))
	q.Insert(13, 0)
	q.Insert(14, 1)
	q.Insert(15, 2)
	v, ok := q.fifo[0].head.data.(int)
	assert.True(t, ok)
	assert.Equal(t, 13, v)
	v, ok = q.fifo[1].head.data.(int)
	assert.True(t, ok)
	assert.Equal(t, 14, v)
	v, ok = q.fifo[2].head.data.(int)
	assert.True(t, ok)
	assert.Equal(t, 15, v)
	assert.Nil(t, q.fifo[0].head.next)
	assert.Nil(t, q.fifo[1].head.next)
	assert.Nil(t, q.fifo[2].head.next)

	q.Insert(23, 0)
	v, ok = q.fifo[0].head.data.(int)
	assert.True(t, ok)
	assert.Equal(t, 13, v)
	v, ok = q.fifo[0].tail.data.(int)
	assert.True(t, ok)
	assert.Equal(t, 23, v)
	assert.Equal(t, q.fifo[0].head.next, q.fifo[0].tail)
	assert.Nil(t, q.fifo[0].tail.next)
}

func TestPQ_Fetch(t *testing.T) {
	q := NewQueuePriorN(uint(3))
	assert.Nil(t, q.Fetch())
	q.Insert(14, 1)
	e := q.Fetch()
	assert.NotNil(t, e)
	v, ok := e.(int)
	assert.True(t, ok)
	assert.Equal(t, 14, v)
}

func TestPQ_Order(t *testing.T) {
	q := NewQueuePriorN(uint(3))
	assert.Nil(t, q.Fetch())
	q.Insert(13, 1)
	q.Insert(4, 2)
	q.Insert(33, 0)
	q.Insert(15, 1)
	e := q.Fetch()
	v, ok := e.(int)
	assert.True(t, ok)
	assert.Equal(t, 33, v)
	e = q.Fetch()
	v, ok = e.(int)
	assert.True(t, ok)
	assert.Equal(t, 13, v)
	e = q.Fetch()
	v, ok = e.(int)
	assert.True(t, ok)
	assert.Equal(t, 15, v)
	e = q.Fetch()
	v, ok = e.(int)
	assert.True(t, ok)
	assert.Equal(t, 4, v)
	assert.Nil(t, q.Fetch())
}

func TestPQ_Concurrency(t *testing.T) {
	q := NewQueuePriorN(uint(3))
	assert.Nil(t, q.Fetch())
	expected := []int{33, 13, 15, 4, 2, 5, 7, 9, 0, 35, 33, 13, 15, 4, 2, 5, 7, 9, 0, 35}
	for i := 0; i < len(expected); i++ {
		go func(i int) {
			q.Insert(expected[i], 0)
		}(i)
	}
	for i := 0; i < len(expected); i++ {
		go func() {
			e := q.Fetch()
			_, ok := e.(int)
			assert.True(t, ok)
		}()
	}
	assert.Nil(t, q.Fetch())
}
