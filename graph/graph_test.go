package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDefaultVertex() Vertex {
	return Vertex{Name: "A name"}
}

func initGraph() Graph {
	graph := NewGraph()
	graph.AddNewVertex("one", createDefaultVertex())
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
			graph := initGraph()
			graph.AddNewVertex("two", createDefaultVertex())

			assert.Len(t, graph.vertices, 2)

			vertex, _ := graph.GetVertex("one")
			assert.Equal(t, createDefaultVertex(), vertex)
		},
	)
	t.Run(
		"test add duplicate vertex name fails",
		func(t *testing.T) {
			graph := initGraph()
			vertex1 := createDefaultVertex()
			vertex1.Name = "Another newer name"
			err := graph.AddNewVertex("one", vertex1)

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
