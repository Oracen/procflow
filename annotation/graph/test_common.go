package graph

import (
	"github.com/Oracen/procflow/core/flags"
)

func init() {
	recordFlow := flags.GetRecordFlow()
	if !recordFlow {
		mockStateManagement()
		StateManager.EnableTrackState()
	}

}
