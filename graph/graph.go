package graph

import "errors"

var ErrGraphVertexAlreadyExists = errors.New("this vertex has already been created under this name")
var ErrGraphVertexNotFound = errors.New("this vertex not found in graph")

type Graph struct {
	vertices VertexCollection
	edges    EdgeCollection
}

func NewGraph() Graph {
	graph := Graph{vertices: VertexCollection{}, edges: EdgeCollection{}}
	return graph
}

func (g *Graph) AddNewVertex(name string, vertex Vertex) (err error) {
	obj, ok := g.vertices[name]
	if ok && (obj.data != vertex.data) {
		return ErrGraphVertexAlreadyExists
	}
	if !ok {
		g.vertices[name] = vertex
	}
	return
}

func (g *Graph) GetVertex(name string) (vertex Vertex, err error) {
	vertex, ok := g.vertices[name]
	if !ok {
		return vertex, ErrGraphVertexNotFound
	}
	return
}

func (g *Graph) AddNewEdge(name string, edge Edge) (err error) {
	obj, ok := g.edges[name]
	if ok && (obj.data != edge.data) {
		return ErrGraphVertexAlreadyExists
	}
	if !ok {
		g.edges[name] = edge
	}
	return
}
