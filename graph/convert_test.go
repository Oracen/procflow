package graph

import (
	"context"
	"testing"

	"github.com/Oracen/procflow/core/constants"
	"github.com/stretchr/testify/assert"
)

func TestConvertToGraphPackage(t *testing.T) {
	t.Run(
		"test node name conversion functions",
		func(t *testing.T) {
			parent := "parent"
			child := "child"
			got := packNames(parent, child)
			want := len(parent) + len(child) + len(constants.StandardDelimiter)
			assert.Len(t, got, want)

			got = unpackNames(parent, got)
			assert.Equal(t, child, got)

		},
	)
	t.Run(
		"test convert simple graph yields dag",
		func(t *testing.T) {
			ctx := context.Background()
			tracker := RegisterTracker(ctx)
			ctx, node := Start(ctx, &tracker, "start", "Start point")
			_, node = Task(ctx, &tracker, []Node{node}, "task", "A task name is longer")
			End(&tracker, []Node{node}, "end", "Endpoint", false)
			tracker.CloseTrace()

			dotGraph := Convert(tracker.Collector.Object.Graph)
			assert.Equal(t, 2, dotGraph.Size())
		},
	)
}
