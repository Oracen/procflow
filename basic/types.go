package basic

import (
	"github.com/Oracen/procflow/core/tracker"
)

type (
	Node    = tracker.Node[string]
	Tracker = tracker.Tracker[string]
)

func RegisterTracker() Tracker {
	return Tracker{}
}
