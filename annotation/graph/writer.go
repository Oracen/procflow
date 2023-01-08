package graph

import (
	"io"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	log "github.com/sirupsen/logrus"
)

func ExportGraphDot(graph graph.Graph[string, string], file io.Writer) {
	// dot -Tsvg -O filename.gv
	if !StateManager.TrackState() {
		return
	}

	err := draw.DOT(graph, file)
	if err != nil {
		log.Error("writing graph dot file failed with error: " + err.Error())
	}

}
