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
			vertices3 := sliceVertices(baseVertices, 2, numVertex)

			// Test on test helper
			assert.Len(t, vertices1, 4)
			assert.Len(t, vertices2, numVertex-4)
			assert.Len(t, vertices3, numVertex-2)

			merged, err := MergeVertices(vertices1, vertices2)
			assert.Nil(t, err)
			assert.Len(t, merged, numVertex)
			assert.Equal(t, baseVertices, merged)

			merged, err = MergeVertices(vertices1, vertices3)
			assert.Nil(t, err)
			assert.Len(t, merged, numVertex)
			assert.Equal(t, baseVertices, merged)
		},
	)

	t.Run(
		"test merge vertices fails if duplicate keys have different data",
		func(t *testing.T) {
			vertices1 := sliceVertices(baseVertices, 0, 4)
			vertices2 := sliceVertices(baseVertices, 3, numVertex)
			vertices2["3"] = vertices2["4"]

			_, err := MergeVertices(vertices1, vertices2)
			assert.ErrorIs(t, err, ErrVertexMergeDuplicate)
		},
	)
}
