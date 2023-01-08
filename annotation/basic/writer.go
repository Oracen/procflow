package basic

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func CreateFile(filename string) io.Writer {
	file, err := os.Create(fmt.Sprintf("%s/graph.gv", filename))
	if err != nil {
		log.Error("graph file creation failed with error: " + err.Error())
	}
	return file
}

func ExportRunTxt(bytes []byte, file io.Writer) {
	// dot -Tsvg -O filename.gv
	if !StateManager.TrackState() {
		return
	}
	_, err := file.Write(bytes)
	if err != nil {
		log.Error("writing graph dot file failed with error: " + err.Error())
	}

}
