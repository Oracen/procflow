package topo_test

import (
	"testing"

	"github.com/Oracen/bpmn-flow/topo"
)

func TestGraphBuildFunctionality(t *testing.T) {
	t.Run(
		"test external package can utilise data",
		func(t *testing.T) {
			graph := topo.NewGraph()
		},
	)
}
