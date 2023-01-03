package topo_test

import (
	"testing"

	"github.com/Oracen/process-flow/topo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type tvertex = topo.Vertex[string]
type tedge = topo.Edge[string]
type tgraph = topo.Graph[string, string]

func createGraph() tgraph {
	return topo.CreateNewGraph[string, string]()
}

func TestPublicApi(t *testing.T) {

	getVertices := func(count int) (testVertices []tvertex) {

		testVertices = []tvertex{}
		for idx := 0; idx < count; idx++ {
			testVertices = append(testVertices, tvertex{uuid.New().String(), uuid.New().String()})
		}
		return
	}

	buildDefaultGraph := func() (graph tgraph, err error) {
		graph = createGraph()
		for _, item := range getVertices(4) {
			name := uuid.New().String()
			err = graph.AddNewVertex(name, item)
			if err != nil {
				return
			}
			for key := range graph.GetAllVertices(true) {
				edge := tedge{name, key, "blah"}
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
