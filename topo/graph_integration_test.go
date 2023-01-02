package topo_test

import (
	"testing"

	"github.com/Oracen/bpmn-flow/topo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGraphBuildFunctionality(t *testing.T) {

	getVertices := func(count int) (testVertices []topo.Vertex) {

		testVertices = []topo.Vertex{}
		for idx := 0; idx < count; idx++ {
			testVertices = append(testVertices, topo.Vertex{uuid.New().String(), topo.VertexData{uuid.New().String()}})
		}
		return
	}

	t.Run(
		"test external package can utilise data and build self loops",
		func(t *testing.T) {
			graph := topo.NewGraph()
			for _, item := range getVertices(4) {
				name := uuid.New().String()
				err := graph.AddNewVertex(name, item)
				assert.Nil(t, err)
				for key := range graph.GetAllVertices(true) {
					edge := topo.Edge{name, key, topo.EdgeData{InvocationName: "blah"}}
					err = graph.AddNewEdge(uuid.New().String(), edge)
					assert.Nil(t, err)
				}
			}
			assert.Len(t, graph.GetAllVertices(false), 4)
			assert.Len(t, graph.GetAllEdges(false), 10)

		},
	)
}
