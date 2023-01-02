package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphBasicFunction(t *testing.T) {
	defaultVertexName := "A name"
	altVertexName := "a random name"

	createDefaultVertex := func(name string) Vertex {
		return Vertex{siteName: name}
	}

	initGraph := func() Graph {
		graph := NewGraph()
		graph.AddNewVertex(defaultVertexName, createDefaultVertex(defaultVertexName))
		return graph
	}

	t.Run(
		"test graph creates correctly",
		func(t *testing.T) {
			graph := NewGraph()
			assert.Len(t, graph.vertices, 0)
		},
	)
	t.Run(
		"test add vertex",
		func(t *testing.T) {
			graph := initGraph()
			assert.Len(t, graph.vertices, 1)

			vertex := createDefaultVertex(altVertexName)
			graph.AddNewVertex(altVertexName, vertex)

			assert.Len(t, graph.vertices, 2)
			got, _ := graph.GetVertex(altVertexName)
			assert.Equal(t, vertex, got)
		},
	)
	t.Run(
		"test add duplicate vertex name fails",
		func(t *testing.T) {
			graph := initGraph()
			vertex := createDefaultVertex(defaultVertexName)

			// Same name but equal data, should not fail
			err := graph.AddNewVertex(defaultVertexName, vertex)
			assert.Equal(t, err, nil)

			// Same name, different data, should fail
			vertex.data = VertexData{taskName: "Get Data"}
			err = graph.AddNewVertex(defaultVertexName, vertex)
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

	t.Run(
		"test add edge",
		func(t *testing.T) {
			graph := initGraph()
			vertex := createDefaultVertex(altVertexName)
			graph.AddNewVertex(altVertexName, vertex)
			assert.Len(t, graph.vertices, 2) // Sanity check

			edge := Edge{defaultVertexName, altVertexName, EdgeData{}}

			graph.AddNewEdge("edge name", edge)
			assert.Len(t, graph.edges, 1)

		},
	)
}
