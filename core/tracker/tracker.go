package tracker

import (
	"sync"

	"github.com/Oracen/procflow/core/collection"
)

type Node[T comparable] struct {
	nodeSite string
	data     T
	previous T
}

type BasicTracker[S any, T comparable] struct {
	traceClosed bool
	collector   collection.Collector[T, S]
	wg          *sync.WaitGroup
}

func RegisterBasicTracker[S any, T comparable](collector collection.Collector[T, S]) BasicTracker[S, T] {
	wg := sync.WaitGroup{}
	return BasicTracker[S, T]{traceClosed: false, collector: collector, wg: &wg}
}

func (t *BasicTracker[S, T]) StartFlow(name string, data T) Node[T] {
	var empty T
	t.collector.AddRelationship(data)
	return Node[T]{name, data, empty}
}

func (t *BasicTracker[S, T]) AddNode(name string, inputs []Node[T], data T) Node[T] {
	var empty T
	return Node[T]{name, data, empty}
}

func (t *BasicTracker[S, T]) EndFlow(name string, inputs []Node[T], data T) {
}

func (t *BasicTracker[S, T]) CloseTrace() {
	t.wg.Wait()
	t.traceClosed = true
}
