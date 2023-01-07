package graph

import (
	"github.com/Oracen/procflow/core/constants"
	"github.com/dominikbraun/graph"
)

func Convert(graphData Graph) (dotGraph graph.Graph[string, string]) {
	dotGraph = graph.New(graph.StringHash, graph.Directed())

	// First instantiate all vertices
	for key, value := range graphData.GetAllVertices(true) {
		name := unpackNames(value.Data.ParentFlow, key)
		dotGraph.AddVertex(
			name,
			graph.VertexAttribute("style", "filled"),
			graph.VertexAttribute("label", value.SiteName),
			graph.VertexAttribute("fillcolor", value.Data.Colour),
			graph.VertexAttribute("shape", value.Data.Shape),
		)
	}

	// Add in primary edges
	for _, value := range graphData.GetAllEdges(true) {
		from, _ := graphData.GetVertex(value.VertexFrom)
		fromName := unpackNames(from.Data.ParentFlow, value.VertexFrom)
		to, _ := graphData.GetVertex(value.VertexTo)
		toName := unpackNames(to.Data.ParentFlow, value.VertexTo)
		dotGraph.AddEdge(
			fromName,
			toName,
			graph.EdgeAttribute("color", value.Data.Colour),
		)

	}

	// Handle nested flows
	for key, value := range graphData.GetAllVertices(true) {
		check := value.Data.ParentFlow == constants.ContextParentDefault ||
			(!value.Data.IsFlowStart && !value.Data.IsFlowEnd)
		if check {
			continue
		}
		name := unpackNames(value.Data.ParentFlow, key)
		start, end := value.Data.ParentFlow, name

		if value.Data.IsFlowEnd {
			start, end = name, value.Data.ParentFlow
		}

		dotGraph.AddEdge(
			start,
			end,
			graph.EdgeAttribute("style", "dotted"),
		)
	}
	return
}
