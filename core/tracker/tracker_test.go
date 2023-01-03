package tracker

import (
	"testing"

	"github.com/Oracen/procflow/core/collection"
)

type (
	utTracker       = Tracker[mockCollectable, string]
	utNodes         = []Node[string]
	mockCollectable = collection.MockCollectable[string]
)

func TestTracker(t *testing.T) {
	t.Run(
		"test tracker properly creates nodes",
		func(t *testing.T) {
			collectable := mockCollectable{Collection: []string{}}
			collector := collection.CreateNewCollector[string, mockCollectable](&collectable)
			tracker := utTracker{collector: collector}
			node1 := tracker.StartFlow("name1", "input")
			node2 := tracker.AddNode("name2", utNodes{node1}, "intermediate")
			tracker.EndFlow("name3", utNodes{node2}, "endpoint")

		},
	)
}
