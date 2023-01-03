package topo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeFunctions(t *testing.T) {

	numEdges := 6
	baseEdges := EdgeCollection[edgeData]{}
	for idx := 0; idx < numEdges; idx++ {
		key := fmt.Sprint(idx)
		baseEdges[key] = Edge[edgeData]{key, key, edgeData{key}}
	}

	sliceEdges := func(edges EdgeCollection[edgeData], start, stop int) (sliced EdgeCollection[edgeData]) {
		sliced = EdgeCollection[edgeData]{}
		for idx := start; idx < stop; idx++ {
			sliced[fmt.Sprint(idx)] = edges[fmt.Sprint(idx)]
		}
		return
	}

	t.Run(
		"test merge vertices",
		func(t *testing.T) {
			slices := []buildSlices{{0, 4}, {4, numEdges}, {2, numEdges}}

			var baseSlice EdgeCollection[edgeData]
			for idx, item := range slices {
				slice := sliceEdges(baseEdges, item.start, item.stop)

				// Quick check on test helper
				assert.Len(t, slice, item.stop-item.start)
				if idx == 0 {
					// Nothing to compare to yet
					baseSlice = slice
					continue
				}

				merged, err := MergeEdges(baseSlice, slice)
				assert.Nil(t, err)
				assert.Len(t, merged, numEdges)
				assert.Equal(t, baseEdges, merged)
			}
		},
	)

	t.Run(
		"test merge vertices fails if duplicate keys have different data",
		func(t *testing.T) {
			edges1 := sliceEdges(baseEdges, 0, 4)
			edges2 := sliceEdges(baseEdges, 3, numEdges)
			edges2["3"] = edges2["4"]

			_, err := MergeEdges(edges1, edges2)
			assert.ErrorIs(t, err, errEdgeMergeDuplicate)
		},
	)
}

type edgeData struct {
	InvocationName string
}
