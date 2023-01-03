package topo_test

import (
	"testing"

	"github.com/Oracen/process-flow/topo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGraphBuildFunctionality(t *testing.T) {

	getVertices := func(count int) (testVertices []topo.Vertex[string]) {

		testVertices = []topo.Vertex[string]{}
		for idx := 0; idx < count; idx++ {
			testVertices = append(testVertices, topo.Vertex[string]{uuid.New().String(), uuid.New().String()})
		}
		return
	}

	buildDefaultGraph := func() (graph topo.Graph[string, string], err error) {
		graph = topo.CreateNewGraph[string, string]()
		for _, item := range getVertices(4) {
			name := uuid.New().String()
			err = graph.AddNewVertex(name, item)
			if err != nil {
				return
			}
			for key := range graph.GetAllVertices(true) {
				edge := topo.Edge[string]{name, key, "blah"}
				err = graph.AddNewEdge(uuid.New().String(), edge)
				if err != nil {
					return
				}
			}
		}
		return
	}

	t.Run(
		"test external package can utilise data and build self-loop edges",
		func(t *testing.T) {
			graph, err := buildDefaultGraph()
			assert.Nil(t, err)
			assert.Len(t, graph.GetAllVertices(false), 4)
			assert.Len(t, graph.GetAllEdges(false), 10)

		},
	)

}
