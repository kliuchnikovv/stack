package stack

import "sync"

// SyncedStack implements concurrency-friendly LIFO data structure.
type SyncedStack[T any] struct {
	Stack[T]
	mutex sync.Mutex
}

// NewSynced - initialises new concurrency-friendly stack.
func NewSynced[T any](capacity int) *SyncedStack[T] {
	return &SyncedStack[T]{
		Stack: *New[T](capacity),
		mutex: sync.Mutex{},
	}
}

// Push - adds a new element to the top of the stack.
func (synced *SyncedStack[T]) Push(data *T) {
	synced.mutex.Lock()
	defer synced.mutex.Unlock()
	synced.Stack.Push(data)
}

// Pop - returns an item from the top of the stack and removes it from the stack.
func (synced *SyncedStack[T]) Pop() (*T, bool) {
	synced.mutex.Lock()
	defer synced.mutex.Unlock()
	return synced.Stack.Pop()
}

// PopN - returns N items (or less) from the top of the stack
// in the order in which they are retrieved (LIFO).
func (synced *SyncedStack[T]) PopN(n int) []*T {
	synced.mutex.Lock()
	defer synced.mutex.Unlock()
	return synced.Stack.PopN(n)
}
