package graph

import "errors"

var (
	ErrGraphVertexAlreadyExists = errors.New("this vertex has already been created under this name")
	ErrGraphVertexNotFound      = errors.New("this vertex not found in graph")
	ErrGraphEdgeAlreadyExists   = errors.New("this edge has already been created under this name")
	ErrGraphEdgeNotFound        = errors.New("this edge not found in graph")
)

type Graph struct {
	vertices VertexCollection
	edges    EdgeCollection
}

func NewGraph() Graph {
	graph := Graph{vertices: VertexCollection{}, edges: EdgeCollection{}}
	return graph
}

func (g *Graph) AddNewVertex(name string, vertex Vertex) (err error) {
	return addGraphItem(g.vertices, name, vertex, ErrGraphVertexAlreadyExists)
}

func (g *Graph) GetVertex(name string) (vertex Vertex, err error) {
	return getGraphItem(g.vertices, name, ErrGraphVertexNotFound)
}

func (g *Graph) AddNewEdge(name string, edge Edge) (err error) {
	for _, vName := range []string{edge.vertexFrom, edge.vertexTo} {
		_, err := g.GetVertex(vName)
		if err != nil {
			return ErrGraphVertexNotFound
		}
	}
	return addGraphItem(g.edges, name, edge, ErrGraphEdgeAlreadyExists)
}

func (g *Graph) GetEdge(name string) (edge Edge, err error) {
	return getGraphItem(g.edges, name, ErrGraphEdgeNotFound)
}

func getGraphItem[T comparable](collection map[string]T, name string, errType error) (item T, err error) {
	item, ok := collection[name]
	if !ok {
		return item, errType
	}
	return

}

func addGraphItem[T comparable](collection map[string]T, name string, item T, errType error) (err error) {
	var value T

	value, ok := collection[name]
	if ok && (value != item) {
		return errType
	}
	if !ok {
		collection[name] = item
	}
	return
}
