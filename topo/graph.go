package topo

import "errors"

var (
	errGraphVertexAlreadyExists = errors.New("this vertex has already been created under this name")
	errGraphVertexNotFound      = errors.New("this vertex not found in graph")
	errGraphEdgeAlreadyExists   = errors.New("this edge has already been created under this name")
	errGraphEdgeNotFound        = errors.New("this edge not found in graph")
)

type Graph[S, T comparable] struct {
	vertices VertexCollection[S]
	edges    EdgeCollection[T]
}

func CreateNewGraph[S, T comparable]() Graph[S, T] {
	graph := Graph[S, T]{vertices: VertexCollection[S]{}, edges: EdgeCollection[T]{}}
	return graph
}

func (g *Graph[S, T]) AddNewVertex(name string, vertex Vertex[S]) (err error) {
	return addGraphItem(g.vertices, name, vertex, errGraphVertexAlreadyExists)
}

func (g *Graph[S, T]) GetVertex(name string) (vertex Vertex[S], err error) {
	return getGraphItem(g.vertices, name, errGraphVertexNotFound)
}

func (g *Graph[S, T]) GetAllVertices(copy bool) (vertices VertexCollection[S]) {
	return getAllItems(g.vertices, copy)
}

func (g *Graph[S, T]) AddNewEdge(name string, edge Edge[T]) (err error) {
	for _, vName := range []string{edge.VertexFrom, edge.VertexTo} {
		_, err := g.GetVertex(vName)
		if err != nil {
			return errGraphVertexNotFound
		}
	}
	return addGraphItem(g.edges, name, edge, errGraphEdgeAlreadyExists)
}

func (g *Graph[S, T]) GetEdge(name string) (edge Edge[T], err error) {
	return getGraphItem(g.edges, name, errGraphEdgeNotFound)
}

func (g *Graph[S, T]) GetAllEdges(copy bool) (edge EdgeCollection[T]) {
	return getAllItems(g.edges, copy)
}

func MergeGraphs[S, T comparable](graph1, graph2 Graph[S, T]) (merged Graph[S, T], err error) {
	vertices, err := MergeVertices(graph1.vertices, graph2.vertices)
	if err != nil {
		return
	}
	edges, err := MergeEdges(graph1.edges, graph2.edges)
	if err != nil {
		return
	}
	merged.vertices = vertices
	merged.edges = edges
	return
}
