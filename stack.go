// Package stack - implementation of stack data structure for Golang
// based on slice data structure.
package stack

type Stack struct {
	buffer []interface{}
	size   int
}

func New(capacity int) *Stack {
	return &Stack{buffer: make([]interface{}, 0, capacity)}
}

// Push - adds a new element to the top of the stack.
func (stack *Stack) Push(value interface{}) {
	if stack.size == len(stack.buffer) {
		stack.buffer = append(stack.buffer, value)
	} else {
		stack.buffer[stack.size] = value
	}
	stack.size++
}

// Peek - returns an item from the top of the stack without removing it.
func (stack *Stack) Peek() (interface{}, bool) {
	if stack.size == 0 {
		return nil, false
	}
	return stack.buffer[stack.lastIndex()], true
}

// Pop - returns an item from the top of the stack and removes it from the stack.
func (stack *Stack) Pop() (interface{}, bool) {
	v, ok := stack.Peek()
	if ok {
		stack.size--
	}
	return v, ok
}

// PopN - returns N items (or less) from the top of the stack in the order in which they are retrieved (LIFO).
func (stack *Stack) PopN(n int) []interface{} {
	var result = make([]interface{}, 0, n)
	if n <= 0 {
		return result
	}
	for i := n; i > 0; i-- {
		item, ok := stack.Pop()
		if !ok {
			return result
		}
		result = append(result, item)
	}
	return result
}

// Size - returns size of stack.
func (stack *Stack) Size() int {
	return stack.size
}

// IsEmpty - returns whether stack is empty or not.
func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

// lastIndex - returns index of last element in the stack's buffer.
func (stack *Stack) lastIndex() int {
	if stack.size == 0 {
		return 0
	}
	return stack.size - 1
}
