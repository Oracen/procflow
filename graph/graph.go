package graph

import "errors"

var ErrGraphVertexAlreadyExists = errors.New("this vertex has already been created under this name")

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
	if ok && (obj == vertex) {
		return ErrGraphVertexAlreadyExists
	}
	g.vertices[name] = vertex
	return
}

func (g *Graph) GetVertex(name string) Vertex {
	return g.vertices[name]
}
