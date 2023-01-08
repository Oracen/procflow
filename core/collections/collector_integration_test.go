package collections_test

import (
	"testing"

	"github.com/Oracen/procflow/core/collections"
	"github.com/stretchr/testify/assert"
)

type MockCol = collections.BasicCollector[int]

func TestPublicApi(t *testing.T) {

	obj := MockCol{Array: []int{-3}}
	col := collections.CreateNewBasicAdapter(&obj)

	for idx := 0; idx < 5000; idx++ {
		value := idx + 1
		go func() {
			err := col.AddRelationship(value)
			assert.Nil(t, err)
		}()

	}

	extra := MockCol{Array: []int{-1, -2}}

	merged, err := col.UnionRelationships(extra)
	assert.Nil(t, err)
	assert.Len(t, merged.Array, 5000+3)
}
