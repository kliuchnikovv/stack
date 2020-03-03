package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackPop(t *testing.T) {
	stack := new(Stack)
	numberOfItems := 10

	for i := 0; i < numberOfItems; i++ {
		stack.Push(i)
	}
	assert.Equal(t, numberOfItems, stack.Len())
	for i, item := range stack.PopN(numberOfItems) {
		assert.Equal(t, (numberOfItems-1)-i, item.(int))
	}
	v, ok := stack.Pop()
	assert.Nil(t, v)
	assert.False(t, ok)
	assert.Equal(t, 0, stack.Len())
}
