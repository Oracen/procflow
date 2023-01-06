package graph

import (
	"github.com/Oracen/procflow/core/topo"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	Errors      = []string
	Constructor = tracker.GraphConstructor[VertexStyle, EdgeStyle]
	Vertex      = topo.Vertex[VertexStyle]
	Edge        = topo.Edge[EdgeStyle]
	Collectable = tracker.GraphCollectable[VertexStyle, EdgeStyle]
	Node        = tracker.Node[Constructor]
	Graph       = topo.Graph[VertexStyle, EdgeStyle]
	Tracker     = tracker.GraphTracker[VertexStyle, EdgeStyle]
)

type VertexStyle struct {
	Colour string
	Shape  string
}

type EdgeStyle struct {
	Colour string
}

func RegisterTracker() Tracker {
	collectable := tracker.CreateNewGraphCollectable[VertexStyle, EdgeStyle]()
	collector := tracker.CreateNewGraphCollector(&collectable)
	return tracker.RegisterGraphTracker(&collector)
}
