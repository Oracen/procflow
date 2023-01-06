package basic

import (
	"github.com/Oracen/procflow/core/collection"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	Collectable = collection.BasicCollectable[string]
	Node        = tracker.Node[string]
	Tracker     = tracker.BasicTracker[string, Collectable]
)

func RegisterTracker() Tracker {
	collectable := Collectable{Collection: []string{}}
	collector := collection.CreateNewCollector[string, Collectable](&collectable)
	return tracker.RegisterBasicTracker(&collector)
}
