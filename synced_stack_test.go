package stack

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
)

func TestPopNFailWhileRoutinesPop(t *testing.T) {
	var startWG, endWG sync.WaitGroup
	var stack = SyncedStack{
		Stack: Stack{
			size: 10,
			buffer: []interface{}{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			},
		},
	}

	for i := 0; i < 1000; i++ {
		result := sync.Map{}

		startWG.Add(5)
		endWG.Add(5)
		for j := 0; j < 5; j++ {
			go func() {
				defer endWG.Done()
				startWG.Done()
				startWG.Wait()

				v, ok := stack.Pop()
				assert.True(t, ok)

				_, ok = result.LoadOrStore(v.(int), true)
				assert.False(t, ok)
			}()
		}

		startWG.Wait()
		values := stack.PopN(5)
		endWG.Wait()

		assert.Less(t, 0, len(values))
		assert.Zero(t, stack.size, values)

		first := values[0].(int)
		for j := 1; j < len(values); j++ {
			assert.Equal(t, first-j, values[j].(int))
		}
	}
}

func BenchmarkSyncedStack_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := NewSynced(numberOfIterations)
		for j := 0; j < numberOfIterations; j++ {
			stack.Push(j)
		}
	}
}

func BenchmarkSyncedStack_Pop(b *testing.B) {
	stack := NewSynced(numberOfIterations)
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

func BenchmarkSyncedStack_PushAndPopRandomly(b *testing.B) {
	stack := NewSynced(numberOfIterations)
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
