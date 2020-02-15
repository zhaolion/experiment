package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCirCleQueue(t *testing.T) {
	size := 3
	queue := NewCirCleQueue(3)
	assert.True(t, queue.IsEmpty())
	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Rear())

	for i := 0; i < size; i++ {
		assert.True(t, queue.Enqueue(1))
	}
	assert.False(t, queue.Enqueue(1))

	for i := 0; i < size; i++ {
		e, ok := queue.Dequeue()
		assert.NotEqual(t, -1, e)
		assert.True(t, ok)
	}
	e, ok := queue.Dequeue()
	assert.Equal(t, -1, e)
	assert.False(t, ok)

	queue = NewCirCleQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Dequeue()
	queue.Enqueue(3)
	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 3, queue.Rear())
	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(4)
	assert.Equal(t, 4, queue.Front())
	assert.Equal(t, 4, queue.Rear())
}
