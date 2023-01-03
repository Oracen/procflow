package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectorCanInitialise(t *testing.T) {

	t.Run(
		"test collector can add and merge nontrivial objects that implement Collectable",
		func(t *testing.T) {
			collectable := mockCollectable{[]int{0}}
			collection := CreateNewCollector[int, mockCollectable](&collectable)
			for idx := 1; idx < 4; idx++ {
				err := collection.AddRelationship(idx)
				assert.Nil(t, err)
			}

			merged, err := collection.UnionRelationships(mockCollectable{[]int{9, 8}})
			assert.Nil(t, err)
			assert.Len(t, merged.collection, 6)
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
