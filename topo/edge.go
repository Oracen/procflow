package topo

import "errors"

var (
	errEdgeMergeDuplicate = errors.New("edges exist in both lists with non-matching data")
)

type Edge[T comparable] struct {
	VertexFrom string
	VertexTo   string
	Data       T
}

type EdgeCollection[T comparable] map[string]Edge[T]

func MergeEdges[T comparable](edges1, edges2 EdgeCollection[T]) (merged EdgeCollection[T], err error) {
	return mergeMapItems(edges1, edges2, errEdgeMergeDuplicate)
}
