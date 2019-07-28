package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_init(t *testing.T) {
	q := &Queue{}
	assert.NotNil(t, q)
	assert.Zero(t, q.Len())
}

func TestQueue_push_one(t *testing.T) {
	q := &Queue{}
	q.Push(11)
	assert.Equal(t, 1, q.Len())

	q.Push(12)
	assert.Equal(t, 2, q.Len())

	q.Push(13)
	q.Push(14)
	q.Push(15)
	assert.Equal(t, 5, q.Len())
}

func TestQueue_push_pop(t *testing.T) {
	q := &Queue{}
	q.Push(11)
	q.Push(12)
	q.Push(13)
	q.Push(14)
	assert.Equal(t, 4, q.Len())

	x, err := q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 11, x)
	assert.Equal(t, 3, q.Len())

	q.Push(42)
	assert.Equal(t, 4, q.Len())

	x, _ = q.Pop()
	assert.Equal(t, 12, x)
	assert.Equal(t, 3, q.Len())

	q.Pop()
	q.Pop()
	assert.Equal(t, 1, q.Len())

	x, _ = q.Pop()
	assert.Equal(t, 42, x)
	assert.Equal(t, 0, q.Len())

	x, err = q.Pop()
	assert.Error(t, ErrEmpty, err)

	x, err = q.Pop()
	assert.Error(t, ErrEmpty, err)
}

func TestQueue_pop_one(t *testing.T) {
	q := &Queue{}
	q.Push(11)
	d, err := q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 11, d)
	assert.Equal(t, 0, q.Len())
}

func TestQueue_push_many(t *testing.T) {
	q := &Queue{}
	for i := 1; i <= 10; i++ {
		q.Push(i)
	}
}
