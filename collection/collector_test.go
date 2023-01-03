package collection

import "testing"

func TestCollectorCanInitialise(t *testing.T) {

	t.Run(
		"test collector can initialise a tracing object",
		func(t *testing.T) {
			collectable := mockCollectable{[]int{}}
			_ = CreateNewCollector[int, mockCollectable](&collectable)
		},
	)

}

type mockCollectable struct {
	collection []int
}

func (m *mockCollectable) AddRelationship(digit int) error {
	m.collection = append(m.collection, digit)
	return nil
}

func (m *mockCollectable) Union(mockCollectable) (mockCollectable, error) {
	return mockCollectable{[]int{}}, nil
}
