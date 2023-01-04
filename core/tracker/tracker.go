package tracker

import (
	"sync"

	"github.com/Oracen/procflow/core/collection"
)

type Node[T comparable] struct {
	nodeSite string
	data     T
}

type Tracker[S any, T comparable] struct {
	traceClosed bool
	collector   collection.Collector[T, S]
	wg          *sync.WaitGroup
}

func RegisterTracker[S, T comparable](collector collection.Collector[T, S]) Tracker[S, T] {
	wg := sync.WaitGroup{}
	return Tracker[S, T]{traceClosed: false, collector: collector, wg: &wg}
}

func (t *Tracker[S, T]) StartFlow(name string, data T) Node[T] {
	return Node[T]{name, data}
}

func (t *Tracker[S, T]) AddNode(name string, inputs []Node[T], data T) Node[T] {
	return Node[T]{name, data}
}

func (t *Tracker[S, T]) EndFlow(name string, inputs []Node[T], data T) {
}

func (t *Tracker[S, T]) CloseTrace() {
	t.wg.Wait()
	t.traceClosed = true
}
