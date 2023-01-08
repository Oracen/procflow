package graph

import (
	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/store"
	"github.com/Oracen/procflow/core/topo"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	// Collector object for shared state
	Constructor = tracker.GraphConstructor[VertexStyle, EdgeStyle]
	// Vertex format required for graphviz constructor
	Vertex = topo.Vertex[VertexStyle]
	// Edge format required for graphviz constructor
	Edge = topo.Edge[EdgeStyle]
	// Collector object for shared state
	Collection = collections.GraphCollector[VertexStyle, EdgeStyle]
	// Tracking data object
	Node = tracker.Node[Constructor]
	// Data storage maps for collector
	Graph = topo.Graph[VertexStyle, EdgeStyle]
	// String node-based flow tracker
	Tracker = tracker.GraphTracker[VertexStyle, EdgeStyle]
	// Shared memory format for graph tracker
	Singleton = store.GlobalState[Collection]
)

// Vertex styling parameters for graph conversion
type VertexStyle struct {
	Colour      string
	Shape       string
	ParentFlow  string
	IsFlowStart bool
}

// Edge styling parameters for graph conversion
type EdgeStyle struct {
	Colour string
}
