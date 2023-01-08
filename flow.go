package procflow

import (
	"github.com/Oracen/procflow/core/store"
)

// Enables the use of global state recording to
func StartFlowRecord(recordFlow bool) {
	if !recordFlow {
		return
	}

	store.StateManager.EnableUseGlobalState()
	store.StateManager.EnableTrackState()
}

func StopFlowRecord(recordFlow bool, filepath string) {
	if !recordFlow {
		return
	}
	store.StateManager.RunExport(filepath)
}
