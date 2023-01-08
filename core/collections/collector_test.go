package collections

import (
	"fmt"
	"sync"
	"testing"

	"github.com/Oracen/procflow/core/topo"
	"github.com/stretchr/testify/assert"
)

func TestCollectorCanInitialise(t *testing.T) {

	t.Run(
		"test basic collector can add and merge",
		func(t *testing.T) {
			nExtra := 5000
			wg := sync.WaitGroup{}
			collector := createNewBasicCollector([]int{0})
			collection := CreateNewBasicAdapter(&collector)
			for idx := 0; idx < nExtra; idx++ {
				wg.Add(1)
				value := idx + 1
				go func() {
					defer wg.Done()
					err := collection.AddRelationship(value)
					assert.Nil(t, err)
				}()

			}
			wg.Wait()
			new := createNewBasicCollector([]int{9, 8})
			merged, err := collection.UnionRelationships(new)
			assert.Nil(t, err)
			assert.Len(t, merged.Array, nExtra+3)
		},
	)

	t.Run(
		"test graph collector can add and merge",
		func(t *testing.T) {
			nExtra := 5000
			wg := sync.WaitGroup{}
			old := topo.CreateNewGraph[int, int]()
			old.AddNewVertex("-2", topo.Vertex[int]{SiteName: "-2", Data: 1})
			old.AddNewVertex("-1", topo.Vertex[int]{SiteName: "-1", Data: 1})
			collector := createNewGraphCollector(old)
			collection := CreateNewGraphAdapter(&collector)
			for idx := 0; idx < nExtra; idx++ {
				wg.Add(1)
				value := idx + 1
				go func() {
					defer wg.Done()
					new := GraphConstructorInner[int, int]{
						VertexName: fmt.Sprint(value),
						Vertex:     topo.Vertex[int]{SiteName: fmt.Sprint(value), Data: 1},
						EdgeName:   "",
						Edge:       topo.Edge[int]{VertexFrom: "-2", VertexTo: "-1", Data: 0},
					}
					err := collection.AddRelationship(new)
					assert.Nil(t, err)
				}()

			}
			wg.Wait()
			new := topo.CreateNewGraph[int, int]()
			new.AddNewVertex("-3", topo.Vertex[int]{SiteName: "-3", Data: 1})
			newCollector := createNewGraphCollector(new)
			merged, err := collection.UnionRelationships(newCollector)
			assert.Nil(t, err)
			assert.Len(t, merged.Graph.GetAllVertices(true), nExtra+3)
		},
	)

}
