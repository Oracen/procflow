package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertexFunctions(t *testing.T) {
	numVertex := 6
	baseVertices := VertexCollection{}
	for idx := 0; idx < numVertex; idx++ {
		baseVertices[fmt.Sprint(idx)] = Vertex{fmt.Sprint(idx), VertexData{fmt.Sprint(idx)}}
	}

	sliceVertices := func(vertices VertexCollection, start, stop int) (sliced VertexCollection) {
		sliced = VertexCollection{}
		for idx := start; idx < stop; idx++ {
			sliced[fmt.Sprint(idx)] = vertices[fmt.Sprint(idx)]
		}
		return
	}

	t.Run(
		"test merge vertices",
		func(t *testing.T) {
			vertices1 := sliceVertices(baseVertices, 0, 4)
			vertices2 := sliceVertices(baseVertices, 4, numVertex)
			assert.Len(t, vertices1, 4)
			assert.Len(t, vertices2, numVertex-4)

			merged := MergeVertices(vertices1, vertices2)
			assert.Len(t, merged, numVertex)
			assert.Equal(t, baseVertices, merged)
		},
	)
}
