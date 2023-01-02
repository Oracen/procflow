package topo

import "errors"

var (
	errEdgeMergeDuplicate = errors.New("edges exist in both lists with non-matching data")
)

type EdgeData struct {
	InvocationName string
}

type Edge struct {
	VertexFrom string
	VertexTo   string
	Data       EdgeData
}

type EdgeCollection map[string]Edge

func MergeEdges(edges1, edges2 EdgeCollection) (merged EdgeCollection, err error) {
	return mergeMapItems(edges1, edges2, errEdgeMergeDuplicate)
}
