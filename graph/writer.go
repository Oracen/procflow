package graph

import (
	"fmt"
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	log "github.com/sirupsen/logrus"
)

func ExportGraphDot(graph graph.Graph[string, string], filename string) {
	// dot -Tsvg -O filename.gv
	file, err := os.Create(fmt.Sprintf("%s.gv", filename))
	if err != nil {
		log.Error("graph file creation failed with error: " + err.Error())
	}
	err = draw.DOT(graph, file)
	if err != nil {
		log.Error("writing graph dot file failed with error: " + err.Error())
	}

}
