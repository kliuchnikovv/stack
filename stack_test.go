package stack

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const numberOfIterations = 100000

func TestStack(t *testing.T) {
	t.Parallel()

	stack := New[int](numberOfIterations)
	for i := 0; i < numberOfIterations; i++ {
		stack.Push(&i)
	}
	assert.Equal(t, numberOfIterations, stack.Size())
	for i, item := range stack.PopN(numberOfIterations) {
		assert.Equal(t, (numberOfIterations-1)-i, *item)
	}
	v, ok := stack.Pop()
	assert.Nil(t, v)
	assert.False(t, ok)
	assert.Zero(t, stack.Size())
}

func TestStack_Push(t *testing.T) {
	t.Parallel()

	stack := new(Stack[int])

	stack.size = 1
	assert.Panics(t, func() { stack.Push(&stack.size) })

	stack.size = -1
	assert.Panics(t, func() { stack.Push(&stack.size) })

	stack.size = 0
	assert.NotPanics(t, func() { stack.Push(&stack.size) })
	assert.Equal(t, 1, stack.size)
	assert.Equal(t, 1, len(stack.buffer))
	assert.Equal(t, 1, *stack.buffer[0])
}

func TestStack_Peek(t *testing.T) {
	t.Parallel()

	stack := new(Stack[int])

	v, ok := stack.Peek()
	assert.False(t, ok)
	assert.Nil(t, v)

	var i = 0
	stack.Push(&i)

	v, ok = stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, 0, *v)
	assert.Equal(t, 1, stack.size)
}

func TestStack_Pop(t *testing.T) {
	t.Parallel()

	stack := new(Stack[int])

	v, ok := stack.Pop()
	assert.False(t, ok)
	assert.Nil(t, v)

	var i = 0
	stack.Push(&i)

	v, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 0, *v)
	assert.Equal(t, 0, stack.size)
}

func TestStack_PopN(t *testing.T) {
	t.Parallel()

	stack := New[int](10)

	items := stack.PopN(10)
	assert.Zero(t, len(items))

	for i := 0; i < 10; i++ {
		stack.Push(&i)
	}

	items = stack.PopN(5)
	assert.Equal(t, 5, len(items))
	for i, item := range items {
		assert.Equal(t, 9-i, *item)
	}

	items = stack.PopN(10)
	assert.Equal(t, 5, len(items))
	for i, item := range items {
		assert.Equal(t, 4-i, *item)
	}
}

func TestStack_Size(t *testing.T) {
	t.Parallel()

	stack := Stack[int]{}
	assert.Zero(t, stack.Size())

	stack.size = 3
	assert.Equal(t, 3, stack.Size())

	stack.size = 0
	assert.Zero(t, stack.Size())

	stack.size = -3
	assert.Equal(t, -3, stack.Size())
}

func TestStack_IsEmpty(t *testing.T) {
	t.Parallel()

	stack := new(Stack[int])
	assert.True(t, stack.IsEmpty())

	stack.size = 3
	assert.False(t, stack.IsEmpty())

	stack.size = 0
	assert.True(t, stack.IsEmpty())

	stack.size = -3
	assert.False(t, stack.IsEmpty())
}

func TestStack_lastIndex(t *testing.T) {
	t.Parallel()

	stack := Stack[int]{}
	assert.Zero(t, stack.lastIndex())

	stack.size = 3
	assert.Equal(t, 2, stack.lastIndex())

	stack.size = 0
	assert.Zero(t, stack.lastIndex())

	stack.size = -3
	assert.Equal(t, -4, stack.lastIndex())
}

func BenchmarkStack_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := New[int](numberOfIterations)
		for j := 0; j < numberOfIterations; j++ {
			stack.Push(&j)
		}
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	stack := New[int](numberOfIterations)
	for i := 0; i < numberOfIterations; i++ {
		stack.Push(&i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		st := stack
		for j := numberOfIterations; j > 0; j-- {
			st.Pop()
		}
	}
}

func BenchmarkStack_PushAndPopRandomly(b *testing.B) {
	stack := New[int](numberOfIterations)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < numberOfIterations; j++ {
			if stack.Size() == 0 || rand.Int() == 0 {
				stack.Push(&j)
			} else {
				stack.Pop()
			}
		}
	}
}
