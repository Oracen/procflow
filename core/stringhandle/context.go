package stringhandle

import (
	"context"

	"github.com/Oracen/procflow/core/constants"
)

func GetParentFlow(ctx context.Context) (parentName string) {
	parentName, ok := ctx.Value(constants.ContextParentFlowKey).(string)
	if parentName == "" || !ok {
		// If no parent, initialise to default parent flow name
		parentName = constants.ContextParentDefault
	}
	return
}
