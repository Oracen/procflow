package graph

import (
	"fmt"
	"io"
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	log "github.com/sirupsen/logrus"
)

func CreateFile(filename string) io.Writer {
	file, err := os.Create(fmt.Sprintf("%s/graph.gv", filename))
	if err != nil {
		log.Error("graph file creation failed with error: " + err.Error())
	}
	return file
}

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
