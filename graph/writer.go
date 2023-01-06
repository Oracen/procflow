package graph

import (
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	log "github.com/sirupsen/logrus"
)

func ExportGraphDot[K comparable, T any](graph graph.Graph[K, T], filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Error("graph file creation failed with error: " + err.Error())
	}
	err = draw.DOT(graph, file)
	if err != nil {
		log.Error("writing graph dot file failed with error: " + err.Error())
	}

}