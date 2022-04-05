package pq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFifo_Create(t *testing.T) {
	q := FIFO{}
	assert.NotNil(t, q)
}

func TestFifo_Pop(t *testing.T) {
	q := FIFO{}
	assert.Nil(t, q.pop())
	q.push(&node{data: 17})
	e := q.pop()
	assert.NotNil(t, e)
	v, ok := e.(int)
	assert.True(t, ok)
	assert.Equal(t, 17, v)
}

func TestFifo_Order(t *testing.T) {
	q := FIFO{}
	q.push(&node{data: 13})
	q.push(&node{data: 2})
	q.push(&node{data: 33})
	expected := []int{13, 2, 33}
	for i := 0; i < 3; i++ {
		e := q.pop()
		v, ok := e.(int)
		assert.True(t, ok)
		assert.Equal(t, expected[i], v)
	}
	assert.Nil(t, q.pop())
}
