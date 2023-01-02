package graph

import "errors"

var (
	ErrVertexMergeDuplicate = errors.New("vertices exist in both lists with non-matching data")
)

type VertexData struct {
	taskName string
}

type Vertex struct {
	siteName string
	data     VertexData
}

type VertexCollection map[string]Vertex

func MergeVertices(vertices1, vertices2 VertexCollection) (merged VertexCollection, err error) {
	merged = VertexCollection{}
	for key, value := range vertices1 {
		merged[key] = value
	}
	for key, value := range vertices2 {
		err = addGraphItem(merged, key, value, ErrVertexMergeDuplicate)
		if err != nil {
			// TODO: Better error communication
			return VertexCollection{}, err
		}
	}
	return
}
