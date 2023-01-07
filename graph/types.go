package graph

import (
	"github.com/Oracen/procflow/core/store"
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
	Singleton   = store.GlobalState[Collectable]
)

type VertexStyle struct {
	Colour      string
	Shape       string
	ParentFlow  string
	IsFlowStart bool
	IsFlowEnd   bool
}

type EdgeStyle struct {
	Colour string
}
