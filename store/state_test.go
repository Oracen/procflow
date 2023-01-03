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
			state := globalState[string]{}
			for idx := range strings {
				state.addObject(&strings[idx])
			}

			// Update basic obecs
			strings[0] = "hello"
			strings[4] = "miracle"

			retrieved := state.getState()
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
			wg := sync.WaitGroup{}

			// Initialise test holders
			var (
				stringSingleton       globalState[string]
				firstState, lastState *globalState[string]
			)

			// Build a lot of backends in goroutines, add data as we do
			for idx := 0; idx < numInstances; idx++ {
				wg.Add(1)
				counter := idx
				go func() {
					state := createGlobalSingleton(&stringSingleton, &lock)

					go func() {
						defer wg.Done()
						s1 := fmt.Sprint(counter)
						s2 := "-" + s1
						state.addObject(&s1)
						state.addObject(&s2)

					}()
					if counter == 0 {
						firstState = state
					}
					if counter == numInstances-1 {
						lastState = state
					}
				}()

			}
			wg.Wait()

			// Check for singleton behaviour
			copyFirst := firstState.getState()
			copyLast := lastState.getState()
			baseState := stringSingleton.getState()

			assert.Len(t, baseState, numInstances*2)
			assert.Equal(t, copyFirst, copyLast)
			assert.Equal(t, copyFirst, baseState)

		},
	)

}
