package basic

import (
	"github.com/Oracen/procflow/core/collection"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	BasicCollection = collection.BasicCollectable[string]
	Node            = tracker.Node[string]
	Tracker         = tracker.BasicTracker[BasicCollection, string]
)

func RegisterTracker() Tracker {
	collectable := BasicCollection{Collection: []string{}}
	collector := collection.CreateNewCollector[string, BasicCollection](&collectable)
	return tracker.RegisterBasicTracker(collector)
}
