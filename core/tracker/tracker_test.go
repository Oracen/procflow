package tracker

import "testing"

type (
	utTracker = Tracker[string]
	utNodes   = []Node[string]
)

func TestTracker(t *testing.T) {
	t.Run(
		"test tracker properly creates nodes",
		func(t *testing.T) {
			tracker := utTracker{}
			node1 := tracker.StartFlow("name1", "input")
			node2 := tracker.AddNode("name2", utNodes{node1}, "intermediate")
			tracker.EndFlow("name3", utNodes{node2}, "endpoint")
		},
	)
}
