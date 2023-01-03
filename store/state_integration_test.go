package store_test

import (
	"sync"
	"testing"

	"github.com/Oracen/process-flow/store"
	"github.com/stretchr/testify/assert"
)

type DefState = store.GlobalState[int]

func TestPublicApi(t *testing.T) {

	numInstances := 3
	lock := sync.Mutex{}
	lockWg := sync.WaitGroup{}

	// Initialise test holders
	var (
		stringSingleton       DefState
		firstState, lastState *DefState
	)

	// Build a lot of backends in goroutines, add data as we do
	for idx := 0; idx < numInstances; idx++ {
		lockWg.Add(1)
		counter := idx
		go func() {
			state := store.CreateGlobalSingleton(&stringSingleton, &lock)
			go func() {
				defer lockWg.Done()
				s1 := counter
				s2 := -1 * counter
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
}
