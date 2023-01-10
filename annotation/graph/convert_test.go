package graph

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/flags"
	"github.com/stretchr/testify/assert"
)

func TestConvertToGraphPackage(t *testing.T) {
	recordFlow := flags.GetRecordFlow()
	if !*recordFlow {
		mockStateManagement()
		StateManager.EnableTrackState()
	}
	t.Run(
		"test convert simple graph yields dag",
		func(t *testing.T) {
			ctx := context.Background()
			tracker := RegisterTracker(ctx)
			ctx, node := Start(&tracker, "start", "Start point")
			_, node = Task(ctx, &tracker, []Node{node}, "task", "A task name is longer")
			End(&tracker, []Node{node}, "end", "Endpoint", true, false)
			tracker.CloseTrace()

			dotGraph := Convert(tracker.Collector.Object.Graph)
			assert.Equal(t, 2, dotGraph.Size())
			_, err := dotGraph.Vertex("start")
			assert.Nil(t, err)
		},
	)

	t.Run(
		"test convert nested graph yields dcg",
		func(t *testing.T) {
			ctx := context.Background()
			tracker := RegisterTracker(ctx)
			ctx, node := Start(&tracker, "start", "Start point")
			ctx, nodeTop := Task(ctx, &tracker, []Node{node}, "task", "A task name is longer")

			tracker2 := RegisterTracker(ctx)
			ctx, node = Start(&tracker2, "startInner", "Start point inner")
			_, node = Task(ctx, &tracker2, []Node{node}, "taskInner", "A task name is inside")
			End(&tracker2, []Node{node}, "endInner", "Endpoint inner", true, false)

			End(&tracker, []Node{nodeTop}, "end", "Endpoint", true, false)
			tracker.CloseTrace()

			g, _ := tracker.Collector.UnionRelationships(*tracker2.Collector.Object)

			dotGraph := Convert(g.Graph)
			assert.Equal(t, 5, dotGraph.Size())
			_, err := dotGraph.Vertex("startInner")
			assert.Nil(t, err)
		},
	)

	t.Run(
		"test exporter writes",
		func(t *testing.T) {
			var mockSingleton Singleton
			collection := collections.CreateNewGraphCollector[VertexStyle, EdgeStyle]()
			buffer := &bytes.Buffer{}

			export := registerGlobal(&mockSingleton, &collection, func(string) io.Writer { return buffer })
			export.ExportRun("file")
			assert.Contains(t, buffer.String(), "strict digraph")
		},
	)
}
