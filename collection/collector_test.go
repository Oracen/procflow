package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectorCanInitialise(t *testing.T) {

	t.Run(
		"test collector can add and merge nontrivial objects that implement Collectable",
		func(t *testing.T) {
			nExtra := 5000
			collectable := MockCollectable{[]int{0}}
			collection := CreateNewCollector[int, MockCollectable](&collectable)
			for idx := 0; idx < nExtra; idx++ {
				value := idx + 1
				go func() {
					err := collection.AddRelationship(value)
					assert.Nil(t, err)
				}()

			}

			merged, err := collection.UnionRelationships(MockCollectable{[]int{9, 8}})
			assert.Nil(t, err)
			assert.Len(t, merged.Collection, nExtra+3)
		},
	)

}
