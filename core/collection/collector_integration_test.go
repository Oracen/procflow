package collection_test

import (
	"testing"

	"github.com/Oracen/process-flow/core/collection"
	"github.com/stretchr/testify/assert"
)

type MockCol = collection.MockCollectable

func TestPublicApi(t *testing.T) {

	obj := MockCol{Collection: []int{-3}}
	col := collection.CreateNewCollector[int, MockCol](&obj)

	for idx := 0; idx < 5000; idx++ {
		value := idx + 1
		go func() {
			err := col.AddRelationship(value)
			assert.Nil(t, err)
		}()

	}

	extra := MockCol{Collection: []int{-1, -2}}

	merged, err := col.UnionRelationships(extra)
	assert.Nil(t, err)
	assert.Len(t, merged.Collection, 5000+3)
}
