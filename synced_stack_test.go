package stack

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopNFailWhileRoutinesPop(t *testing.T) {
	var startWG, endWG sync.WaitGroup
	var stack = Stack[int]{}

	for i := 1; i <= 10; i++ {
		stack.Push(&i)
	}

	for i := 0; i < 1000; i++ {
		result := sync.Map{}
		st := SyncedStack[int]{
			Stack: stack,
			mutex: sync.Mutex{},
		}
		startWG.Add(5)
		endWG.Add(5)
		for j := 0; j < 5; j++ {
			go func() {
				defer endWG.Done()
				startWG.Done()
				startWG.Wait()

				v, ok := st.Pop()
				assert.True(t, ok)
				_, ok = result.LoadOrStore(v, true)
				assert.False(t, ok)
			}()
		}

		startWG.Wait()
		values := st.PopN(5)
		endWG.Wait()

		assert.Less(t, 0, len(values))
		assert.Zero(t, st.size, values)

		for j := 1; j < len(values); j++ {
			assert.Equal(t, *values[0]-j, *values[j])
		}
	}
}

func BenchmarkSyncedStack_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := NewSynced[int](numberOfIterations)
		for j := 0; j < numberOfIterations; j++ {
			stack.Push(&j)
		}
	}
}

func BenchmarkSyncedStack_Pop(b *testing.B) {
	stack := NewSynced[int](numberOfIterations)
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

func BenchmarkSyncedStack_PushAndPopRandomly(b *testing.B) {
	stack := NewSynced[int](numberOfIterations)
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
