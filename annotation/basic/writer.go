package basic

import (
	"io"

	log "github.com/sirupsen/logrus"
)

// Function for handling the text write operation
func ExportRunTxt(bytes []byte, file io.Writer) {
	if !StateManager.TrackState() {
		return
	}
	_, err := file.Write(bytes)
	if err != nil {
		log.Error("writing graph dot file failed with error: " + err.Error())
	}

}
