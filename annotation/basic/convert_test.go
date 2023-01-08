package basic

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/constants"
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
			ctx, node := Start(ctx, &tracker, "start", "Start point")
			_, node = Task(ctx, &tracker, []Node{node}, "task", "A task name is longer")
			End(&tracker, []Node{node}, "end", "Endpoint", false)
			tracker.CloseTrace()

			bytes := Convert(tracker.Collector.Object.Array)
			assert.Equal(t, 56, len(bytes))
			want := fmt.Sprintf(
				"%s%staskInner%s%s",
				taskLabel.TASK,
				constants.StandardDelimiter,
				constants.StandardDelimiter,
				taskLabel.TASK,
			)
			assert.NotContains(t, tracker.Collector.Object.Array, want)
		},
	)

	t.Run(
		"test convert nested graph yields dcg",
		func(t *testing.T) {
			ctx := context.Background()
			tracker := RegisterTracker(ctx)
			ctx, node := Start(ctx, &tracker, "start", "Start point")
			ctx, nodeTop := Task(ctx, &tracker, []Node{node}, "task", "A task name is longer")

			tracker2 := RegisterTracker(ctx)
			ctx, node = Start(ctx, &tracker2, "startInner", "Start point inner")
			_, node = Task(ctx, &tracker2, []Node{node}, "taskInner", "A task name is inside")
			End(&tracker2, []Node{node}, "endInner", "Endpoint inner", false)

			End(&tracker, []Node{nodeTop}, "end", "Endpoint", false)
			tracker.CloseTrace()

			g, _ := tracker.Collector.UnionRelationships(*tracker2.Collector.Object)

			bytes := Convert(g.Array)
			assert.Equal(t, 122, len(bytes))

			want := fmt.Sprintf(
				"%s%staskInner%s%s",
				taskLabel.TASK,
				constants.StandardDelimiter,
				constants.StandardDelimiter,
				taskLabel.TASK,
			)
			assert.Contains(t, g.Array, want)
		},
	)

	t.Run(
		"test exporter writes",
		func(t *testing.T) {
			var mockSingleton Singleton
			collection := collections.CreateNewBasicCollector[string]()
			collection.Array = append(collection.Array, []string{"one", "two"}...)
			buffer := &bytes.Buffer{}

			export := registerGlobal(&mockSingleton, &collection, func(string) io.Writer { return buffer })
			export.ExportRun("file")
			assert.Contains(t, buffer.String(), "\n")
		},
	)
}
