package tracker

import (
	"sync"
	"testing"

	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/topo"
	"github.com/stretchr/testify/assert"
)

type (
	utBasicTracker     = BasicTracker[string]
	utBasicNodes       = []Node[string]
	utGraphNode        = Node[utGraphConstructor]
	utGraphConstructor = GraphConstructor[string, string]
	utGraphCollection  = collections.GraphCollector[string, string]
	utVertex           = topo.Vertex[string]
	utEdge             = topo.Edge[string]
	mockCollection     = collections.BasicCollector[string]
	utGraph            = topo.Graph[string, string]
)

func TestTracker(t *testing.T) {
	t.Run(
		"test basic tracker creates nodes",
		func(t *testing.T) {
			collection := mockCollection{Array: []string{}}
			tracker := RegisterBasicTracker(&collection, "parent")

			node1 := tracker.StartFlow("input")
			node2 := tracker.AddNode(utBasicNodes{node1}, "intermediate")
			tracker.EndFlow(utBasicNodes{node2}, "endpoint")

			tracker.CloseTrace()
			assert.True(t, tracker.traceClosed)
			assert.Len(t, collection.Array, 3)
		},
	)

	t.Run(
		"test trackers properly implement interface",
		func(t *testing.T) {
			trivial := func(t Tracker[string]) bool {
				return t.CloseTrace()
			}

			wgDummy := sync.WaitGroup{}
			collection := mockCollection{Array: []string{}}
			basicTracker := RegisterBasicTracker(&collection, "parent")

			assert.True(t, trivial(&basicTracker))

			trivialGraph := func(t Tracker[utGraphConstructor]) bool {
				return t.CloseTrace()
			}
			graphCollection := utGraphCollection{Graph: utGraph{}, Wg: &wgDummy}
			graphTracker := RegisterGraphTracker(&graphCollection, "parent")

			assert.True(t, trivialGraph(&graphTracker))

		},
	)
}

func TestGraphTracker(t *testing.T) {
	t.Run(
		"test graph tracker creates nodes",
		func(t *testing.T) {
			collection := collections.CreateNewGraphCollector[string, string]()
			tracker := RegisterGraphTracker(&collection, "parent")

			collection.AddTask()
			startData1 := utGraphConstructor{
				"start 1",
				utVertex{SiteName: "start name", Data: "no-data"},
				""}
			startData2 := utGraphConstructor{
				"start 2",
				utVertex{SiteName: "start name 2", Data: "ditto-data"},
				""}
			node1a := tracker.StartFlow(startData1)
			node1b := tracker.StartFlow(startData2)

			intermediateData := utGraphConstructor{
				"intermediate",
				utVertex{SiteName: "intermediate name", Data: "some-data"},
				"nil by",
			}
			node2 := tracker.AddNode([]utGraphNode{node1a, node1b}, intermediateData)

			endData := utGraphConstructor{
				"end",
				utVertex{SiteName: "end name", Data: "all-data"},
				"shows over",
			}
			tracker.EndFlow([]utGraphNode{node2}, endData)

			tracker.CloseTrace()
			assert.True(t, tracker.traceClosed)
			assert.Len(t, collection.Graph.GetAllVertices(true), 4)
			assert.Len(t, collection.Graph.GetAllEdges(true), 3)
		},
	)

	t.Run(
		"test graph tracker reports errors",
		func(t *testing.T) {
			collection := collections.CreateNewGraphCollector[string, string]()
			tracker := RegisterGraphTracker(&collection, "parent")
			collection.AddTask()
			startData1 := utGraphConstructor{
				"start 1",
				utVertex{SiteName: "start name", Data: "no-data"},
				""}
			startData2 := utGraphConstructor{
				"start 1",
				utVertex{SiteName: "start name", Data: "changed-data"},
				""}
			node1a := tracker.StartFlow(startData1)
			node1b := tracker.StartFlow(startData2)

			intermediateData := utGraphConstructor{
				"intermediate",
				utVertex{SiteName: "intermediate name", Data: "some-data"},
				"nil by",
			}
			tracker.AddNode([]utGraphNode{node1a, node1b}, intermediateData)

			tracker.CloseTrace()
			assert.True(t, tracker.traceClosed)
			assert.Len(t, collection.Graph.GetAllVertices(true), 1)
			assert.Len(t, collection.Errors, 1)
		},
	)
}
