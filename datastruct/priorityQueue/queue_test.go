package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Push(t *testing.T) {
	queue := NewPriorityQueue(100)

	queue.Push(Node{0, 1})
	assert.Equal(t, 0, queue.Top().value)

	queue.Push(Node{3, 1})
	assert.Equal(t, 0, queue.Top().value)

	queue.Push(Node{3, 2})
	assert.Equal(t, 3, queue.Top().value)

	queue.Push(Node{6, 6})
	assert.Equal(t, 6, queue.Top().value)

	queue.Push(Node{12, 5})
	assert.Equal(t, 6, queue.Top().value)

	queue.Push(Node{13, 8})
	assert.Equal(t, 13, queue.Top().value)
}

func Test_PushPop(t *testing.T) {
	queue := NewPriorityQueue(100)

	queue.Push(Node{0, 1})
	queue.Push(Node{3, 1})
	queue.Push(Node{3, 2})
	queue.Push(Node{6, 6})
	queue.Push(Node{12, 5})
	queue.Push(Node{13, 8})
	assert.Equal(t, 13, queue.Top().value)

	assert.Equal(t, 13, queue.Pop().value)
	assert.Equal(t, 6, queue.Pop().value)
	assert.Equal(t, 12, queue.Top().value)
	assert.Equal(t, 12, queue.Pop().value)

	queue.Push(Node{24, 8})
	assert.Equal(t, 24, queue.Top().value)
}
