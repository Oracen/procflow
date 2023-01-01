package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			graph := NewGraph()
			vertex0 := Vertex{}
			vertex1 := Vertex{}
			graph.AddNewVertex("one", vertex0)
			graph.AddNewVertex("two", vertex1)

			assert.Len(t, graph.vertices, 2)

			assert.Equal(t, vertex0, graph.GetVertex("one"))
		},
	)
	t.Run(
		"test add duplicate vertex name fails",
		func(t *testing.T) {
			graph := NewGraph()
			vertex0 := Vertex{}
			vertex1 := Vertex{}
			graph.AddNewVertex("one", vertex0)
			err := graph.AddNewVertex("one", vertex1)

			assert.ErrorIs(t, err, ErrGraphVertexAlreadyExists)
		},
	)
}
