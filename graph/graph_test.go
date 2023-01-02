package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	defaultVertexName     = "A name"
	altVertexName         = "a random name"
	defaultEdgeName       = "An Edge name"
	altEdgeName           = "An Edge nom de gurre"
	nonexistentVertexName = "Cest ne pas vertex"
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
			assert.Nil(t, err)

			// Same name, different data, should fail
			vertex.data = VertexData{taskName: "Get Data"}
			err = graph.AddNewVertex(defaultVertexName, vertex)
			assert.ErrorIs(t, err, errGraphVertexAlreadyExists)
		},
	)

	t.Run(
		"test get missing vertex name fails",
		func(t *testing.T) {
			graph := initGraph()
			_, err := graph.GetVertex("two")

			assert.ErrorIs(t, err, errGraphVertexNotFound)
		},
	)
}

func TestGraphEdgeFunction(t *testing.T) {
	type edgeBuild struct {
		name string
		edge Edge
	}

	createEdgePair := func(name1, name2 string) []edgeBuild {
		return []edgeBuild{
			{defaultEdgeName, Edge{name1, name2, EdgeData{}}},
			{altEdgeName, Edge{name2, name1, EdgeData{}}},
		}
	}
	t.Run(
		"test add and get edge",
		func(t *testing.T) {
			graph := createBasicVertices()
			assert.Len(t, graph.vertices, 2) // Sanity check

			edges := createEdgePair(defaultVertexName, altVertexName)
			assert.Len(t, graph.edges, 0) // Sanity check
			for idx, eBuild := range edges {
				err := graph.AddNewEdge(eBuild.name, eBuild.edge)
				assert.Nil(t, err)
				assert.Len(t, graph.edges, idx+1)
				got, _ := graph.GetEdge(eBuild.name)
				assert.Equal(t, eBuild.edge, got)
			}
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
			assert.Nil(t, err)

			// Same name, different data, should fail
			edge.data = EdgeData{invocationName: "Get Data"}
			err = graph.AddNewEdge(defaultEdgeName, edge)
			assert.ErrorIs(t, err, errGraphEdgeAlreadyExists)
		},
	)
	t.Run(
		"test add edge fails if both vertices don't exist",
		func(t *testing.T) {
			graph := createBasicVertices()

			edges := createEdgePair(defaultVertexName, nonexistentVertexName)
			for _, eBuild := range edges {
				err := graph.AddNewEdge(eBuild.name, eBuild.edge)
				assert.ErrorIs(t, err, errGraphVertexNotFound)
				assert.NotErrorIs(t, err, errGraphEdgeNotFound)
				assert.Len(t, graph.edges, 0)
			}

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
