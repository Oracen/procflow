package topo

import "errors"

var (
	errVertexMergeDuplicate = errors.New("vertices exist in both lists with non-matching data")
)

type VertexData struct {
	TaskName string
}

type Vertex struct {
	SiteName string
	Data     VertexData
}

type VertexCollection map[string]Vertex

func MergeVertices(vertices1, vertices2 VertexCollection) (merged VertexCollection, err error) {

	merged = VertexCollection{}
	for key, value := range vertices1 {
		// Create copy for safety
		merged[key] = value
	}

	for key, value := range vertices2 {
		err = addGraphItem(merged, key, value, errVertexMergeDuplicate)
		if err != nil {
			// TODO: Better error communication
			return VertexCollection{}, err
		}
	}
	return
}
