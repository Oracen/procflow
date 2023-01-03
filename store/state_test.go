package store

import (
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

			// Ensure pointers are maintained
			assert.Equal(t, strings, retrieved)

		},
	)

}
