package basic

import (
	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	Collection = collections.BasicCollector[string]
	Node       = tracker.Node[string]
	Tracker    = tracker.BasicTracker[string]
)

func RegisterTracker() Tracker {
	collection := Collection{Array: []string{}}
	return tracker.RegisterBasicTracker(&collection)
}
