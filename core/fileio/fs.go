// Generic utilities for interacting with the filesystem
package fileio

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func CreateFileEncapsulation(filename string) func(string) io.Writer {
	return func(filepath string) io.Writer {
		filepath = strings.TrimSuffix(filepath, "/")
		file, err := os.Create(fmt.Sprintf("%s/%s", filepath, filename))
		if err != nil {
			log.Error("graph file creation failed with error: " + err.Error())
		}
		return file
	}

}
