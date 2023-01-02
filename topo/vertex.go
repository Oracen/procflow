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
	return mergeMapItems(vertices1, vertices2, errVertexMergeDuplicate)
}
