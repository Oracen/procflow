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
			collectable := mockCollectable{[]int{0}}
			collection := CreateNewCollector[int, mockCollectable](&collectable)
			for idx := 0; idx < nExtra; idx++ {
				value := idx + 1
				go func() {
					err := collection.AddRelationship(value)
					assert.Nil(t, err)
				}()

			}

			merged, err := collection.UnionRelationships(mockCollectable{[]int{9, 8}})
			assert.Nil(t, err)
			assert.Len(t, merged.collection, nExtra+3)
		},
	)

}

type mockCollectable struct {
	collection []int
}

func (m *mockCollectable) Add(digit int) error {
	m.collection = append(m.collection, digit)
	return nil
}

func (m *mockCollectable) Union(other mockCollectable) (mockCollectable, error) {
	collection := append(m.collection, other.collection...)
	return mockCollectable{collection}, nil
}
