package basic

import (
	"github.com/Oracen/procflow/core/collection"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	BasicCollection = collection.MockCollectable[string]
	Node            = tracker.Node[string]
	Tracker         = tracker.Tracker[BasicCollection, string]
)

func RegisterTracker() Tracker {
	return Tracker{}
}
