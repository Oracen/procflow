package topo

import "errors"

var (
	errGraphVertexAlreadyExists = errors.New("this vertex has already been created under this name")
	errGraphVertexNotFound      = errors.New("this vertex not found in graph")
	errGraphEdgeAlreadyExists   = errors.New("this edge has already been created under this name")
	errGraphEdgeNotFound        = errors.New("this edge not found in graph")
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
	return addGraphItem(g.vertices, name, vertex, errGraphVertexAlreadyExists)
}

func (g *Graph) GetVertex(name string) (vertex Vertex, err error) {
	return getGraphItem(g.vertices, name, errGraphVertexNotFound)
}

func (g *Graph) GetAllVertices(copy bool) (vertices VertexCollection) {
	return getAllItems(g.vertices, copy)
}

func (g *Graph) AddNewEdge(name string, edge Edge) (err error) {
	for _, vName := range []string{edge.VertexFrom, edge.VertexTo} {
		_, err := g.GetVertex(vName)
		if err != nil {
			return errGraphVertexNotFound
		}
	}
	return addGraphItem(g.edges, name, edge, errGraphEdgeAlreadyExists)
}

func (g *Graph) GetEdge(name string) (edge Edge, err error) {
	return getGraphItem(g.edges, name, errGraphEdgeNotFound)
}

func (g *Graph) GetAllEdges(copy bool) (edge EdgeCollection) {
	return getAllItems(g.edges, copy)
}

func MergeGraphs(graph1, graph2 Graph) (merged Graph, err error) {
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
