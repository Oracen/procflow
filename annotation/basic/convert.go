package basic

import (
	"io"
	"log"
	"strings"
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

	converted := Convert(collected.Array)
	fileHandle := e.fsHandler(filename)
	ExportRunTxt(converted, fileHandle)
}

// Conversion function to map from string array to exportable bytes format
func Convert(stringData Array) (bytes []byte) {
	joined := strings.Join(stringData, "\n")
	return []byte(joined)
}
