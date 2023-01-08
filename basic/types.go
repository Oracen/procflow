package basic

import (
	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/store"
	"github.com/Oracen/procflow/core/tracker"
)

type (
	Errors     = []string
	Collection = collections.BasicCollector[string]
	Node       = tracker.Node[string]
	Graph      = []string
	Tracker    = tracker.BasicTracker[string]
	Singleton  = store.GlobalState[Collection]
)
