package topo

import "errors"

var (
	errVertexMergeDuplicate = errors.New("vertices exist in both lists with non-matching data")
)

type Vertex[T comparable] struct {
	SiteName string
	Data     T
}

type VertexCollection[T comparable] map[string]Vertex[T]

func MergeVertices[T comparable](vertices1, vertices2 VertexCollection[T]) (merged VertexCollection[T], err error) {
	return mergeMapItems(vertices1, vertices2, errVertexMergeDuplicate)
}
