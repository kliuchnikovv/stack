package stack

type Stack struct {
	array []interface{}
	size  int
}

func (stack *Stack) lastIndex() int {
	if stack.IsEmpty() {
		return 0
	}
	return stack.size - 1
}

func (stack *Stack) Push(value interface{}) {
	if stack.size == len(stack.array) {
		stack.array = append(stack.array, value)
	} else {
		stack.array[stack.size] = value
	}
	stack.size++
}

func (stack *Stack) Peek() (interface{}, bool) {
	if stack.IsEmpty() {
		return nil, false
	}
	return stack.array[stack.lastIndex()], true
}

func (stack *Stack) Pop() (interface{}, bool) {
	v, ok := stack.Peek()
	if ok {
		// Delete last element.
		stack.array[stack.lastIndex()] = nil
		stack.size--
	}
	return v, ok
}

func (stack *Stack) PopN(n int) []interface{} {
	var result = make([]interface{}, 0)
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

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}
