package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDefaultVertex() Vertex {
	return Vertex{SiteName: "A name"}
}

func initGraph() Graph {
	graph := NewGraph()
	graph.AddNewVertex(createDefaultVertex())
	return graph
}

func TestGraph(t *testing.T) {
	t.Run(
		"test graph creates correctly",
		func(t *testing.T) {
			graph := NewGraph()
			assert.Len(t, graph.vertices, 0)
		},
	)
	t.Run(
		"test add vertices",
		func(t *testing.T) {
			name := "a random name"
			graph := initGraph()
			vertex := createDefaultVertex()
			vertex.SiteName = name
			graph.AddNewVertex(vertex)

			assert.Len(t, graph.vertices, 2)

			got, _ := graph.GetVertex(name)
			assert.Equal(t, vertex, got)
		},
	)
	t.Run(
		"test add duplicate vertex name fails",
		func(t *testing.T) {
			graph := initGraph()
			vertex1 := createDefaultVertex()
			vertex1.Data = VertexData{ActivityName: "Get Data"}
			err := graph.AddNewVertex(vertex1)

			assert.ErrorIs(t, err, ErrGraphVertexAlreadyExists)
		},
	)

	t.Run(
		"test get missing vertex name fails",
		func(t *testing.T) {
			graph := initGraph()
			_, err := graph.GetVertex("two")

			assert.ErrorIs(t, err, ErrGraphVertexNotFound)
		},
	)
}
