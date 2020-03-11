package stack

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

// TODO: make tests great again

const numberOfIterations = 1000000

func TestStackPop(t *testing.T) {
	stack := new(Stack)

	for i := 0; i < numberOfIterations; i++ {
		stack.Push(i)
	}
	assert.Equal(t, numberOfIterations, stack.Size())
	for i, item := range stack.PopN(numberOfIterations) {
		assert.Equal(t, (numberOfIterations-1)-i, item.(int))
	}
	v, ok := stack.Pop()
	assert.Nil(t, v)
	assert.False(t, ok)
	assert.Equal(t, 0, stack.Size())
}

func BenchmarkStack_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := Stack{}
		for j := 0; j < numberOfIterations; j++ {
			stack.Push(j)
		}
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	stack := Stack{}
	for i := 0; i < numberOfIterations; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		st := stack
		for j := numberOfIterations; j > 0; j-- {
			st.Pop()
		}
	}
}

func BenchmarkPushPopRandomly(b *testing.B) {
	stack := Stack{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < numberOfIterations; j++ {
			if stack.Size() == 0 || rand.Int() == 0 {
				stack.Push(j)
			} else {
				stack.Pop()
			}
		}
	}
}
