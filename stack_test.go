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

func BenchmarkStack_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := Stack{}
		for i := 0; i < 1000000; i++ {
			stack.Push(i)
		}
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	b.StopTimer()
	stack := Stack{}
	for i := 0; i < 1000000; i++ {
		stack.Push(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		st := stack
		for i := 1000000; i > 0; i-- {
			st.Pop()
		}
	}
}
