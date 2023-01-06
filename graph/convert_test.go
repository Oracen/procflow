package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToGraphPackage(t *testing.T) {
	t.Run(
		"test convert simple graph yields dag",
		func(t *testing.T) {
			tracker := RegisterTracker()
			node := Start(&tracker, "start", "Start point")
			node = Task(&tracker, []Node{node}, "task", "A task name is longer")
			End(&tracker, []Node{node}, "end", "Endpoint", false)
			tracker.CloseTrace()

			dotGraph := Convert(tracker.Collector.Object.Graph)
			assert.Equal(t, 2, dotGraph.Size())
		},
	)
}
