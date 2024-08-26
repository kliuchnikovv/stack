// Package stack - implementation of stack data structure for Golang
// based on slice data structure.
package stack

// Stack implements simple LIFO data structure.
// If concurrent access from multiple locations
// needed, use SyncedStack.
type Stack[T any] struct {
	buffer []*T
	size   int
}

// New - initialises new stack with given capacity.
func New[T any](capacity int) *Stack[T] {
	if capacity < 0 {
		capacity = 0
	}
	return &Stack[T]{buffer: make([]*T, 0, capacity)}
}

// Push - adds a new element to the top of the stack.
func (stack *Stack[T]) Push(value *T) {
	if stack.size == len(stack.buffer) {
		stack.buffer = append(stack.buffer, value)
	} else {
		stack.buffer[stack.size] = value
	}
	stack.size++
}

// Peek - returns an item from the top of the stack without removing it.
func (stack *Stack[T]) Peek() (*T, bool) {
	if stack.size == 0 {
		return nil, false
	}
	return stack.buffer[stack.lastIndex()], true
}

// Pop - returns an item from the top of the stack and removes it from the stack.
func (stack *Stack[T]) Pop() (*T, bool) {
	v, ok := stack.Peek()
	if ok {
		stack.size--
	}

	return v, ok
}

// PopN - returns N items (or less) from the top of the stack
// in the order in which they are retrieved (LIFO).
func (stack *Stack[T]) PopN(n int) []*T {
	var result = make([]*T, 0, n)
	if n <= 0 {
		return result
	}

	for i := 0; i < n && stack.size > 0; i++ {
		result = append(result, stack.buffer[stack.lastIndex()])
		stack.size--
	}

	return result
}

// Size - returns size of stack.
func (stack *Stack[T]) Size() int {
	return stack.size
}

// IsEmpty - returns whether stack is empty or not.
func (stack *Stack[T]) IsEmpty() bool {
	return stack.size == 0
}

// ToSlice - returns slice representation of stack
func (stack *Stack[T]) ToSlice() []*T {
	var result = make([]*T, stack.size)
	for i := 0; i < len(result); i++ {
		result[i] = stack.buffer[i]
	}

	return result
}

// lastIndex - returns index of last element in the stack's buffer.
func (stack *Stack[T]) lastIndex() int {
	if stack.size == 0 {
		return 0
	}

	return stack.size - 1
}
