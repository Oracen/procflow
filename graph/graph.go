package graph

import "errors"

var ErrGraphVertexAlreadyExists = errors.New("this vertex has already been created under this name")
var ErrGraphVertexNotFound = errors.New("this vertex not found in graph")

type Graph struct {
	vertices GraphVertices
}

type GraphVertices map[string]Vertex

func NewGraph() Graph {
	graph := Graph{vertices: GraphVertices{}}
	return graph
}

func (g *Graph) AddNewVertex(name string, vertex Vertex) (err error) {
	obj, ok := g.vertices[name]
	if ok && (obj != vertex) {
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
