package graph

import (
	"github.com/dominikbraun/graph"
)

func Convert(graphData Graph) (dotGraph graph.Graph[string, string]) {
	dotGraph = graph.New(graph.StringHash, graph.Directed())
	for key, value := range graphData.GetAllVertices(true) {

		dotGraph.AddVertex(
			key,
			graph.VertexAttribute("label", value.SiteName),
			graph.VertexAttribute("color", value.Data.Colour),
			graph.VertexAttribute("shape", value.Data.Shape),
		)
	}
	for _, value := range graphData.GetAllEdges(true) {
		dotGraph.AddEdge(
			value.VertexFrom,
			value.VertexTo,
			graph.EdgeAttribute("color", value.Data.Colour),
		)
	}
	return
}
