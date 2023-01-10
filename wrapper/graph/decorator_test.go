package basic

import (
	"context"
	"errors"
	"testing"

	"github.com/Oracen/procflow/annotation/graph"
	"github.com/Oracen/procflow/core/tracker"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	insideStr string
	insideInt int
}

func TestTrackerCtx(t *testing.T) {
	successPayload := testStruct{"success", 1}
	createTestTracker := func() *TrackerCtx {
		ctx := context.Background()
		decorator := CreateNewTrackerCtx(ctx)
		return &decorator
	}

	tFunc1 := func(ctx context.Context, str string) (int, error) {
		if str != "success" {
			return 0, errors.New("Inner error")
		}
		return 1, nil
	}

	tFunc2 := func(ctx context.Context, i int) (testStruct, error) {
		if i != 1 {
			return testStruct{"fail", 0}, errors.New("Inner error")
		}
		return successPayload, nil
	}

	tFunc3 := func(ctx context.Context, payload testStruct) (int, error) {
		if payload != successPayload {
			return 0, errors.New("Inner error")
		}
		return 1, nil
	}

	t.Run(
		"test basic tracker creates nodes",
		func(t *testing.T) {
			decorator := createTestTracker()
			assert.NotNil(t, decorator)
		},
	)

	t.Run(
		"test basic node passing works",
		func(t *testing.T) {
			dec := createTestTracker()

			start := "success"

			msg1, err := Start(dec, "start1", "startDesc1", tFunc1, start)
			assert.Nil(t, err)

			msg2, err := Task(dec, "task2", "taskDesc2", tFunc2, msg1)
			assert.Nil(t, err)
			_, err = End(dec, "task3", "taskDesc3", tFunc3, msg2)
			assert.Nil(t, err)

			dec.CloseTrace()

		},
	)

	t.Run(
		"test node packing and unpacking",
		func(t *testing.T) {
			input := "payload"
			nRepeats := 4

			nodes1, nodes2, nodes3 := []graph.Node{}, []graph.Node{}, []graph.Node{}
			params := graph.Constructor{
				Name:     "",
				Vertex:   graph.TaskVertex("", ""),
				EdgeData: graph.StandardEdge(),
			}
			data := tracker.ConstructGraphNode(params)
			for idx := 0; idx < nRepeats; idx++ {
				nodes1 = append(nodes1, data)
				nodes2 = append(nodes2, data)
				nodes3 = append(nodes3, data)
			}

			msg1 := RepackMessage(input, nodes1, nodes2, nodes3)

			payload, nodes := UnpackMessage(msg1)
			assert.Equal(t, input, payload)
			assert.Len(t, nodes, nRepeats*3)
		},
	)
}
