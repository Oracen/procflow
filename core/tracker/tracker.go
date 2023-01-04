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

type Tracker[T comparable] interface {
	StartFlow(name string, data T) Node[T]
	AddNode(name string, inputs []Node[T], data T) Node[T]
	EndFlow(name string, inputs []Node[T], data T)
	CloseTrace() bool
}

type BasicTracker[S comparable, T any] struct {
	traceClosed bool
	collector   *collection.Collector[S, T]
	wg          *sync.WaitGroup
}

func RegisterBasicTracker[S comparable, T any](collector *collection.Collector[S, T]) BasicTracker[S, T] {
	wg := sync.WaitGroup{}
	return BasicTracker[S, T]{traceClosed: false, collector: collector, wg: &wg}
}

func (t *BasicTracker[S, T]) StartFlow(name string, data S) Node[S] {
	var empty S
	t.collector.AddRelationship(data)
	return Node[S]{name, data, empty}
}

func (t *BasicTracker[S, T]) AddNode(name string, inputs []Node[S], data S) Node[S] {
	var empty S
	return Node[S]{name, data, empty}
}

func (t *BasicTracker[S, T]) EndFlow(name string, inputs []Node[S], data S) {
}

func (t *BasicTracker[S, T]) CloseTrace() bool {
	t.wg.Wait()
	t.traceClosed = true
	return t.traceClosed
}

func RegisterGraphTracker[S comparable, T any](collector *collection.Collector[S, T]) {

}
