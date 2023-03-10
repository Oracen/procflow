package graph

import (
	"io"
	"log"

	"github.com/Oracen/procflow/core/constants"
	"github.com/Oracen/procflow/core/stringhandle"
	"github.com/dominikbraun/graph"
)

type exporter struct {
	singleton *Singleton
	fsHandler func(string) io.Writer
}

func (e *exporter) ExportRun(filename string) {
	var collected Collection

	outputs := e.singleton.GetState()
	for idx, item := range outputs {
		item.WaitForFinish()
		if idx == 0 {
			collected = item
			continue
		}
		merged, err := collected.Union(item)
		if err != nil {
			log.Fatal("Fatal error in procflow.graph export")
		}
		collected = merged

	}
	dotGraph := Convert(collected.Graph)
	fileHandle := e.fsHandler(filename)
	ExportGraphDot(dotGraph, fileHandle)
}

// Conversion function to map from topo.graph to exportable graphviz format
func Convert(graphData Graph) (dotGraph graph.Graph[string, string]) {
	dotGraph = graph.New(graph.StringHash, graph.Directed())

	// First instantiate all vertices
	for key, value := range graphData.GetAllVertices(true) {
		name := stringhandle.UnpackNames(value.Data.ParentFlow, key)
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
		fromName := stringhandle.UnpackNames(from.Data.ParentFlow, value.VertexFrom)
		to, _ := graphData.GetVertex(value.VertexTo)
		toName := stringhandle.UnpackNames(to.Data.ParentFlow, value.VertexTo)
		dotGraph.AddEdge(
			fromName,
			toName,
			graph.EdgeAttribute("color", value.Data.Colour),
		)

	}

	// Handle nested flows
	for key, value := range graphData.GetAllVertices(true) {
		check := value.Data.ParentFlow == constants.ContextParentDefault || !value.Data.IsFlowStart
		if check {
			continue
		}
		name := stringhandle.UnpackNames(value.Data.ParentFlow, key)
		start, end := value.Data.ParentFlow, name

		dotGraph.AddEdge(
			start,
			end,
			graph.EdgeAttribute("style", "dotted"),
		)
	}
	return
}
