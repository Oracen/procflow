package collection

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectorCanInitialise(t *testing.T) {

	t.Run(
		"test collector can add and merge nontrivial objects that implement Collectable",
		func(t *testing.T) {
			nExtra := 5000
			wg := sync.WaitGroup{}
			collectable := MockCollectable[int]{[]int{0}}
			collection := CreateNewCollector[int, MockCollectable[int]](&collectable)
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
			merged, err := collection.UnionRelationships(MockCollectable[int]{[]int{9, 8}})
			assert.Nil(t, err)
			assert.Len(t, merged.Collection, nExtra+3)
		},
	)

}
