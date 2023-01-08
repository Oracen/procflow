package basic

import (
	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/store"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	// Collector object for shared state
	Collection = collections.BasicCollector[string]
	// Tracking data object
	Node = tracker.Node[string]
	// Data storage array for collector
	Array = []string
	// String node-based flow tracker
	Tracker = tracker.BasicTracker[string]
	// Shared memory format for basic tracker
	Singleton = store.GlobalState[Collection]
)
