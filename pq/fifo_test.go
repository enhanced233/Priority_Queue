package pq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestFifo_Push(t *testing.T) {
	q := FIFO{}
	q.push(&node{data: 17})
	v, ok := q.head.data.(int)
	assert.True(t, ok)
	assert.Equal(t, 17, v)
	assert.Equal(t, q.head, q.tail)
	assert.Nil(t, q.tail.next)

	q.push(&node{data: 55})
	v, ok = q.head.data.(int)
	assert.True(t, ok)
	assert.Equal(t, 17, v)
	v, ok = q.tail.data.(int)
	assert.True(t, ok)
	assert.Equal(t, 55, v)
	assert.Equal(t, q.head.next, q.tail)
	assert.Nil(t, q.tail.next)
}

func TestFifo_Order(t *testing.T) {
	q := FIFO{}
	q.push(&node{data: 13})
	q.push(&node{data: 2})
	e := q.pop()
	v, ok := e.(int)
	assert.True(t, ok)
	assert.Equal(t, 13, v)
	e = q.pop()
	v, ok = e.(int)
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	assert.Nil(t, q.pop())
}
