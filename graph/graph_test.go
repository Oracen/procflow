package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	defaultVertexName = "A name"
	altVertexName     = "a random name"
	defaultEdgeName   = "An Edge name"
	altEdgeName       = "An Edge nom de gurre"
)

func TestGraphBasicFunction(t *testing.T) {

	t.Run(
		"test graph creates correctly",
		func(t *testing.T) {
			graph := NewGraph()
			assert.Len(t, graph.vertices, 0)
		},
	)
}
func TestGraphVertexFunction(t *testing.T) {
	t.Run(
		"test add and get vertex",
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
}

func TestGraphEdgeFunction(t *testing.T) {
	t.Run(
		"test add and get edge",
		func(t *testing.T) {
			graph := createBasicVertices()
			assert.Len(t, graph.vertices, 2) // Sanity check

			edge1 := Edge{defaultVertexName, altVertexName, EdgeData{}}
			edge2 := Edge{altVertexName, defaultEdgeName, EdgeData{}}

			graph.AddNewEdge(defaultEdgeName, edge1)
			assert.Len(t, graph.edges, 1)
			graph.AddNewEdge(altEdgeName, edge2)
			assert.Len(t, graph.edges, 2)

			got, _ := graph.GetEdge(defaultEdgeName)
			assert.Equal(t, edge1, got)
		},
	)

	t.Run(
		"test add duplicate edge name fails",
		func(t *testing.T) {
			graph := createBasicVertices()

			edge := Edge{defaultVertexName, altVertexName, EdgeData{}}
			graph.AddNewEdge(defaultEdgeName, edge)
			assert.Len(t, graph.edges, 1) // Sanity check

			// Same name but equal data, should not fail
			err := graph.AddNewEdge(defaultEdgeName, edge)
			assert.Equal(t, err, nil)

			// Same name, different data, should fail
			edge.data = EdgeData{invocationName: "Get Data"}
			err = graph.AddNewEdge(defaultEdgeName, edge)
			assert.ErrorIs(t, err, ErrGraphEdgeAlreadyExists)
		},
	)
}

func createDefaultVertex(name string) Vertex {
	return Vertex{siteName: name}
}

func initGraph() Graph {
	graph := NewGraph()
	graph.AddNewVertex(defaultVertexName, createDefaultVertex(defaultVertexName))
	return graph
}

func createBasicVertices() Graph {
	graph := initGraph()
	vertex := createDefaultVertex(altVertexName)
	graph.AddNewVertex(altVertexName, vertex)
	return graph
}
