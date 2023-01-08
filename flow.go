package procflow

import (
	"github.com/Oracen/procflow/core/store"
)

// Enables the use of global state recording to trace processes within the program
func StartFlowRecord(recordFlow bool) {
	if !recordFlow {
		return
	}

	store.StateManager.EnableUseGlobalState()
	store.StateManager.EnableTrackState()
}

// Awaits the end of the trackers then exports all tracked flows to disk
func StopFlowRecord(recordFlow bool, filepath string) {
	if !recordFlow {
		return
	}
	store.StateManager.RunExport(filepath)
}
