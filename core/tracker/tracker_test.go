package tracker

import (
	"sync"
	"testing"

	"github.com/Oracen/procflow/core/collection"
	"github.com/stretchr/testify/assert"
)

type (
	utTracker       = BasicTracker[string, mockCollectable]
	utNodes         = []Node[string]
	mockCollectable = collection.BasicCollectable[string]
)

func TestTracker(t *testing.T) {
	t.Run(
		"test basic tracker properly creates nodes",
		func(t *testing.T) {
			wg := sync.WaitGroup{}

			collectable := mockCollectable{Collection: []string{}}
			collector := collection.CreateNewCollector[string, mockCollectable](&collectable)
			tracker := utTracker{traceClosed: true, collector: &collector, wg: &wg}
			node1 := tracker.StartFlow("name1", "input")
			node2 := tracker.AddNode("name2", utNodes{node1}, "intermediate")
			tracker.EndFlow("name3", utNodes{node2}, "endpoint")

			tracker.CloseTrace()
			assert.True(t, tracker.traceClosed)
		},
	)

	t.Run(
		"test trackers properly implement interface",
		func(t *testing.T) {
			trivial := func(t Tracker[string]) bool {
				return t.CloseTrace()
			}

			collectable := mockCollectable{Collection: []string{}}
			collector := collection.CreateNewCollector[string, mockCollectable](&collectable)
			basicTracker := RegisterBasicTracker(&collector)

			assert.True(t, trivial(&basicTracker))

			// graphTracker := RegisterGraphTracker(&collector)

			// assert.True(t, trivial(&basicTracker))

		},
	)
}
