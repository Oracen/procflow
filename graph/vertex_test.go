package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type buildSlices struct {
	start, stop int
}

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
			slices := []buildSlices{{0, 4}, {4, numVertex}, {2, numVertex}}

			var baseSlice VertexCollection
			for idx, item := range slices {
				slice := sliceVertices(baseVertices, item.start, item.stop)

				// Quick check on test helper
				assert.Len(t, slice, item.stop-item.start)
				if idx == 0 {
					// Nothing to compare to yet
					baseSlice = slice
					continue
				}

				merged, err := MergeVertices(baseSlice, slice)
				assert.Nil(t, err)
				assert.Len(t, merged, numVertex)
				assert.Equal(t, baseVertices, merged)
			}
		},
	)

	t.Run(
		"test merge vertices fails if duplicate keys have different data",
		func(t *testing.T) {
			vertices1 := sliceVertices(baseVertices, 0, 4)
			vertices2 := sliceVertices(baseVertices, 3, numVertex)
			vertices2["3"] = vertices2["4"]

			_, err := MergeVertices(vertices1, vertices2)
			assert.ErrorIs(t, err, errVertexMergeDuplicate)
		},
	)
}
