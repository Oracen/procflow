package procflow

import (
	"github.com/Oracen/procflow/core/store"
)

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
