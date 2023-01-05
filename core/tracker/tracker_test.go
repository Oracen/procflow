package tracker

import (
	"sync"
	"testing"

	"github.com/Oracen/procflow/core/collection"
	"github.com/Oracen/procflow/core/topo"
	"github.com/stretchr/testify/assert"
)

type (
	utBasicTracker     = BasicTracker[string, mockCollectable]
	utBasicNodes       = []Node[string]
	utGraphNode        = Node[utGraphConstructor]
	utGraphConstructor = GraphConstructor[string, string]
	utGraphCollectable = GraphCollectable[string, string]
	utVertex           = topo.Vertex[string]
	utEdge             = topo.Edge[string]
	mockCollectable    = collection.BasicCollectable[string]
	utGraph            = topo.Graph[string, string]
)

func TestTracker(t *testing.T) {
	t.Run(
		"test basic tracker creates nodes",
		func(t *testing.T) {
			wg := sync.WaitGroup{}

			collectable := mockCollectable{Collection: []string{}}
			collector := collection.CreateNewCollector[string, mockCollectable](&collectable)
			tracker := utBasicTracker{traceClosed: true, collector: &collector, wg: &wg}
			node1 := tracker.StartFlow("input")
			node2 := tracker.AddNode(utBasicNodes{node1}, "intermediate")
			tracker.EndFlow(utBasicNodes{node2}, "endpoint")

			tracker.CloseTrace()
			assert.True(t, tracker.traceClosed)
			assert.Len(t, collectable.Collection, 3)
		},
	)

	t.Run(
		"test trackers properly implement interface",
		func(t *testing.T) {
			trivial := func(t Tracker[string]) bool {
				return t.CloseTrace()
			}

			collectable := mockCollectable{Collection: []string{}}
			basicCollector := collection.CreateNewCollector[string, mockCollectable](&collectable)
			basicTracker := RegisterBasicTracker(&basicCollector)

			assert.True(t, trivial(&basicTracker))

			trivialGraph := func(t Tracker[utGraphConstructor]) bool {
				return t.CloseTrace()
			}
			graphCollectable := utGraphCollectable{Graph: utGraph{}}
			graphCollector := CreateNewGraphCollector(&graphCollectable)
			graphTracker := RegisterGraphTracker(&graphCollector)

			assert.True(t, trivialGraph(&graphTracker))

		},
	)
}

func TestGraphTracker(t *testing.T) {
	t.Run(
		"test graph tracker creates nodes",
		func(t *testing.T) {
			wg := sync.WaitGroup{}

			collectable := CreateNewGraphCollectable[string, string]()
			collector := CreateNewGraphCollector(&collectable)
			tracker := GraphTracker[string, string]{traceClosed: false, collector: &collector, wg: &wg}

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
			assert.Len(t, collectable.Graph.GetAllVertices(true), 4)
			assert.Len(t, collectable.Graph.GetAllEdges(true), 3)
		},
	)

	t.Run(
		"test graph tracker reports errors",
		func(t *testing.T) {
			wg := sync.WaitGroup{}

			collectable := CreateNewGraphCollectable[string, string]()
			collector := CreateNewGraphCollector(&collectable)
			tracker := GraphTracker[string, string]{traceClosed: false, collector: &collector, wg: &wg}

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
			assert.Len(t, collectable.Graph.GetAllVertices(true), 1)
			assert.Len(t, collectable.Errors, 1)
		},
	)
}
