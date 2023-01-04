package basic

import (
	"github.com/Oracen/procflow/core/collection"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	BasicCollectable = collection.BasicCollectable[string]
	Node             = tracker.Node[string]
	Tracker          = tracker.BasicTracker[string, BasicCollectable]
)

func RegisterTracker() Tracker {
	collectable := BasicCollectable{Collection: []string{}}
	collector := collection.CreateNewCollector[string, BasicCollectable](&collectable)
	return tracker.RegisterBasicTracker(&collector)
}
