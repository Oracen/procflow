package stringhandle

import (
	"fmt"
	"strings"

	"github.com/Oracen/procflow/core/constants"
)

func PackNames(parentNodeName, currentNodeName string) string {
	return fmt.Sprintf("%s%s%s", parentNodeName, constants.StandardDelimiter, currentNodeName)
}

func UnpackNames(parentNodeName, packedName string) string {
	prefix := fmt.Sprintf("%s%s", parentNodeName, constants.StandardDelimiter)
	return strings.TrimPrefix(packedName, prefix)
}

func StartFlowName(currentNodeName string) string {
	return fmt.Sprintf("%s %s", constants.StartNodeDescriptionPrefix, currentNodeName)
}

func EndFlowName(currentNodeName string) string {
	return fmt.Sprintf("%s %s", constants.StartNodeDescriptionPrefix, currentNodeName)
}
