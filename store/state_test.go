package store

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateSystems(t *testing.T) {
	t.Run(
		"test state behaviour",
		func(t *testing.T) {
			strings := []string{
				"goodbye",
				"world",
				"what",
				"a",
				"tragedy",
				"it",
				"is",
			}
			state := createEmptyGlobalState[string]()
			for idx := range strings {
				state.AddObject(&strings[idx])
			}

			// Update basic obecs
			strings[0] = "hello"
			strings[4] = "miracle"

			retrieved := state.GetState()
			// Check lengths same
			assert.Len(t, retrieved, len(strings))

			// Ensure pointers maintain values...
			assert.Equal(t, strings, retrieved)

			// ...but are copies of the underlying object
			for idx := range strings {
				assert.NotSame(t, &strings[idx], &retrieved[idx])
			}

		},
	)
	t.Run(
		"test singleton forms properly",
		func(t *testing.T) {
			// Init constancs
			numInstances := 5000
			lock := sync.Mutex{}
			lockWg := sync.WaitGroup{}

			// Initialise test holders
			var (
				stringSingleton       GlobalState[string]
				firstState, lastState *GlobalState[string]
			)

			// Build a lot of backends in goroutines, add data as we do
			for idx := 0; idx < numInstances; idx++ {
				lockWg.Add(1)
				counter := idx
				go func() {
					state := CreateGlobalSingleton(&stringSingleton, &lock)
					go func() {
						defer lockWg.Done()
						s1 := fmt.Sprint(counter)
						s2 := "-" + s1
						state.AddObject(&s1)
						state.AddObject(&s2)
					}()
					if counter == 0 {
						firstState = state
					}
					if counter == numInstances-1 {
						lastState = state
					}
				}()

			}
			lockWg.Wait()

			// Check for singleton behaviour
			copyFirst := firstState.GetState()
			copyLast := lastState.GetState()
			baseState := stringSingleton.GetState()

			assert.Len(t, baseState, numInstances*2)
			assert.Equal(t, copyFirst, copyLast)
			assert.Equal(t, copyFirst, baseState)

		},
	)

}
